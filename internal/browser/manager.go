// internal/browser/manager.go
package browser

import (
	"context"

	"github.com/evesfect/LLLMapi/internal/utils/logger"
)

type Manager struct {
	logger logger.Logger
}

func NewManager(cfg interface{}, log logger.Logger) *Manager {
	return &Manager{
		logger: log,
	}
}

func (m *Manager) Start(ctx context.Context) error {
	return nil
}

func (m *Manager) Stop(ctx context.Context) error {
	return nil
}
