package routes

import (
	"github.com/Perazzojoao/WebSocket-ChatApp/src/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type ChiRouter struct {
	router    *chi.Mux
	wsHandler *handlers.WsHandler
}

func NewChiRouter(ws *handlers.WsHandler) *ChiRouter {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	router := &ChiRouter{
		router:    r,
		wsHandler: ws,
	}

	router.Routes()

	return router
}

func (c *ChiRouter) GetRouter() *chi.Mux {
	return c.router
}

func (c *ChiRouter) Routes() {
	c.router.Get("/ws", c.wsHandler.WsHand)
	c.router.Get("/broadcast", c.wsHandler.Broadcast)
}
