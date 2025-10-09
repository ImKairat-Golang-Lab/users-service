package http

import (
	"github.com/go-chi/chi/v5"
)


type HttpHandler struct {
	mux *chi.Mux
	userHandler UserHandler
}

func NewHttpHandler(mux *chi.Mux, handler UserHandler) *HttpHandler {
	return &HttpHandler{mux: mux, userHandler: handler}
}

func (h *HttpHandler) Handle() {
	h.mux.Group(func(r chi.Router) {
		r.Post("/register", h.userHandler.UserRegister)
	})
}
