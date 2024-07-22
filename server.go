package todo

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		MaxHeaderBytes: 1 << 20,          // 1 mb
		ReadTimeout:    10 * time.Second, // ожиданиe read request-a от клиента
		WriteTimeout:   10 * time.Second, // ожиданиe write request-a от клиента
	}

	return s.httpServer.ListenAndServe() // запускает бесконечный цикл, который слушает запросы для дальнейшей обработки
}

// обертка для httpServer.Shutdown
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
