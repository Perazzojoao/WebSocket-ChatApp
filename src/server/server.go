package server

import (
	"flag"
	"log"
	"net/http"

	"github.com/Perazzojoao/WebSocket-ChatApp/src/handlers"
	"github.com/Perazzojoao/WebSocket-ChatApp/src/routes"
	"github.com/Perazzojoao/WebSocket-ChatApp/src/server/websocket"
	"github.com/go-chi/chi/v5"
)

var (
	url = flag.String("url", "localhost:3000", "The url to listen on")
)

type Server struct {
	router *chi.Mux
}

func NewServer() *Server {
	ws := websocket.NewWsServer()
	wsHandler := handlers.NewWsHandler(ws)
	r := routes.NewChiRouter(wsHandler)
	return &Server{
		router: r.GetRouter(),
	}
}

func (s *Server) Start() {
	flag.Parse()

	srv := http.Server{
		Addr:    *url,
		Handler: s.router,
	}

	log.Printf("Server listening on http://%s", *url)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
