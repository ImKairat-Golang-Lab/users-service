package http

import (
	"fmt"
	"net"
	"net/http"

	ports "github.com/ImKairat-Golang-Lab/users-service/internal/ports"
)

type HTTPServer struct {
	mux         *http.ServeMux
	userHandler UserHandler
	logger ports.Logger
}

func NewHTTPServer(handler UserHandler, logger ports.Logger) *HTTPServer {
	mux := http.NewServeMux()
	s := &HTTPServer{
		mux: mux,
		userHandler: handler,
		logger: logger,
	}
	s.registerRoutes()

	return s
}

func (s *HTTPServer) Start(ip net.IP, port int) error {
	addr := fmt.Sprintf("%s:%d", ip.String(), port)

	s.logger.Info("starting HTTP server", map[string]any{"addr": addr})
	if err := http.ListenAndServe(addr, s.mux); err != nil {
		s.logger.Error("server failed", map[string]any{"error": err})
		return err
	}

	return nil
}

func (s *HTTPServer) registerRoutes() {
	s.mux.HandleFunc("/register", s.userHandler.UserRegister)
}