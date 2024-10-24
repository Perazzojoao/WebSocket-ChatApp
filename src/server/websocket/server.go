package websocket

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type connMap map[*websocket.Conn]bool

type WsServer struct {
	upgrader *websocket.Upgrader
	clients  connMap
}

func NewWsServer(config ...websocket.Upgrader) *WsServer {
	var upgrader websocket.Upgrader
	if len(config) > 0 {
		upgrader = config[0]
	} else {
		upgrader = websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}
	}
	return &WsServer{
		upgrader: &upgrader,
		clients:  make(connMap),
	}
}

func (w *WsServer) Upgrade(wr http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	return w.upgrader.Upgrade(wr, r, nil)
}

func (w *WsServer) HandleConn(ws *websocket.Conn) {
	w.clients[ws] = true
	defer ws.Close()

	for {
		// Read in a message
		_, msg, err := ws.ReadMessage()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println(err)
			break
		}
		// Print out that message for clarity
		log.Printf("Message received: %s", string(msg))
		// Write out that message
		err = ws.WriteMessage(websocket.TextMessage, []byte("Message received"))
		if err != nil {
			log.Println(err)
			break
		}
	}
}

func (w *WsServer) Broadcast(msg string, conn ...*websocket.Conn) {
	var myConn *websocket.Conn = nil
	if len(conn) > 0 {
		myConn = conn[0]
	}

	for client := range w.clients {
		if myConn != nil && client == myConn {
			continue
		}
		go func(client *websocket.Conn) {
			err := client.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				log.Println(err)
				client.Close()
				delete(w.clients, client)
			}
		}(client)
	}
}

func (w *WsServer) HandleBroadcast(ws *websocket.Conn) {
	w.clients[ws] = true
	defer ws.Close()

	for {
		// Read in a message
		_, msg, err := ws.ReadMessage()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println(err)
			break
		}
		// Print out that message for clarity
		log.Printf("Message received: %s", string(msg))

		w.Broadcast(string(msg), ws)
	}
}
