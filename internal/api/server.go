// internal/api/server.go
package api

import (
	"context"
	"net/http"

	"github.com/evesfect/LLLMapi/internal/queue"
	"github.com/evesfect/LLLMapi/internal/utils/logger"
)

type Server struct {
	server *http.Server
	queue  *queue.Manager
	logger logger.Logger
}

func NewServer(cfg interface{}, qm *queue.Manager, log logger.Logger) *Server {
	return &Server{
		queue:  qm,
		logger: log,
	}
}

func (s *Server) Start() error {
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	return nil
}
