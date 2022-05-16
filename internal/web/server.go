package web

import (
	"net/http"

	"golang.org/x/net/context"
)

type Server struct {
	server *http.Server
}

func (s *Server) ServerRun(port string, handler http.Handler) error {
	s.server = &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}

	return s.server.ListenAndServe()
}

func (s *Server) ShutDown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
