package http

import (
	"context"
	"github.com/mereiamangeldin/One-lab-Homework-1/config"
	"github.com/mereiamangeldin/One-lab-Homework-1/transport/http/handler"
	"net/http"
)

type Server struct {
	cfg        *config.Config
	httpServer *http.Server
	handler    *handler.Manager
}

func NewServer(cfg *config.Config, handler *handler.Manager) *Server {
	return &Server{cfg: cfg, handler: handler}
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) ShutDown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
