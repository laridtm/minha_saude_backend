package http

import (
	"context"
	"log"
	"net/http"
	"time"
)

const (
	serverReadTimeout     = 5 * time.Second
	serverWriteTimeout    = 55 * time.Second
	serverShutdownTimeout = 60 * time.Second
)

type server struct {
	srv *http.Server
}

func NewServer(port string, handler http.Handler) *server {
	return &server{
		srv: &http.Server{
			Addr:         ":" + port,
			Handler:      handler,
			ReadTimeout:  serverReadTimeout,
			WriteTimeout: serverWriteTimeout,
		},
	}
}

func (s *server) Start() {
	go func() {
		if err := s.srv.ListenAndServe(); err != nil {
			log.Printf("error on start server: %s", err.Error())
		}
	}()
}

func (s *server) ShutDown() {
	log.Printf("shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), serverShutdownTimeout)
	defer cancel()

	err := s.srv.Shutdown(ctx)
	if err != nil && err != http.ErrServerClosed {
		log.Printf("could not shutdown in 60s: %v", err)
	}

	log.Printf("server gracefully stopped")
}
