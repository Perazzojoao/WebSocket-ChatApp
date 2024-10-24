package handlers

import (
	"log"
	"net/http"

	"github.com/Perazzojoao/WebSocket-ChatApp/src/server/websocket"
)

type WsHandler struct {
	ws *websocket.WsServer
}

func NewWsHandler(ws *websocket.WsServer) *WsHandler {
	return &WsHandler{
		ws: ws,
	}
}

func (h *WsHandler) WsHand(w http.ResponseWriter, r *http.Request) {
	conn, err := h.ws.Upgrade(w, r)
	if err != nil {
		log.Println(err)
		return
	}
	go h.ws.HandleConn(conn)
}

func (h *WsHandler) Broadcast(w http.ResponseWriter, r *http.Request) {
	conn, err := h.ws.Upgrade(w, r)
	if err != nil {
		log.Println(err)
		return
	}
	go h.ws.HandleBroadcast(conn)
}
