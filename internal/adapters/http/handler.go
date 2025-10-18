package http

import (
	"fmt"
	"net"
	"net/http"
)

type HTTPServer struct {
	mux         *http.ServeMux
	userHandler UserHandler
}

func NewHTTPServer(handler UserHandler) *HTTPServer {
	mux := http.NewServeMux()
	s := &HTTPServer{
		mux:         mux,
		userHandler: handler,
	}
	s.registerRoutes()

	return s
}

func (s *HTTPServer) Start(ip net.IP, port int) error {
	addr := fmt.Sprintf("%s:%d", ip.String(), port)

	s.userHandler.logger.Info("starting HTTP server", map[string]any{"addr": addr})
	if err := http.ListenAndServe(addr, s.mux); err != nil {
		s.userHandler.logger.Error("server failed", map[string]any{"error": err})
		return err
	}

	return nil
}

func (s *HTTPServer) registerRoutes() {
	s.mux.HandleFunc("/register", s.userHandler.UserRegister)
}
