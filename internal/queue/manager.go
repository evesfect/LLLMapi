// internal/queue/manager.go
package queue

import (
	"context"

	"github.com/evesfect/LLLMapi/internal/browser"
	"github.com/evesfect/LLLMapi/internal/utils/logger"
)

type Manager struct {
	browser *browser.Manager
	logger  logger.Logger
}

func NewManager(cfg interface{}, bm *browser.Manager, log logger.Logger) *Manager {
	return &Manager{
		browser: bm,
		logger:  log,
	}
}

func (m *Manager) Start(ctx context.Context) error {
	return nil
}

func (m *Manager) Stop(ctx context.Context) error {
	return nil
}
