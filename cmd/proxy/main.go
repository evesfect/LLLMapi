package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/evesfect/LLLMapi/internal/api"
	"github.com/evesfect/LLLMapi/internal/browser"
	"github.com/evesfect/LLLMapi/internal/config"
	"github.com/evesfect/LLLMapi/internal/queue"
	"github.com/evesfect/LLLMapi/internal/utils/logger"
)

func main() {
	// Initialize context with cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Initialize logger
	log := logger.NewLogger()

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	// Initialize components
	browserManager := browser.NewManager(cfg.Browser, log)
	queueManager := queue.NewManager(cfg.Queue, browserManager, log)
	apiServer := api.NewServer(cfg.API, queueManager, log)

	// Start all services
	if err := startServices(ctx, browserManager, queueManager, apiServer, log); err != nil {
		log.Fatal("Failed to start services:", err)
	}

	// Handle graceful shutdown
	handleShutdown(ctx, cancel, browserManager, queueManager, apiServer, log)
}

func startServices(
	ctx context.Context,
	browserManager *browser.Manager,
	queueManager *queue.Manager,
	apiServer *api.Server,
	log logger.Logger,
) error {
	if err := browserManager.Start(ctx); err != nil {
		return fmt.Errorf("failed to start browser manager: %w", err)
	}
	log.Info("Browser manager started successfully")

	if err := queueManager.Start(ctx); err != nil {
		return fmt.Errorf("failed to start queue manager: %w", err)
	}
	log.Info("Queue manager started successfully")

	if err := apiServer.Start(); err != nil {
		return fmt.Errorf("failed to start API server: %w", err)
	}
	log.Info("API server started successfully")

	return nil
}

func handleShutdown(
	ctx context.Context,
	cancel context.CancelFunc,
	browserManager *browser.Manager,
	queueManager *queue.Manager,
	apiServer *api.Server,
	log logger.Logger,
) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan
	log.Info("Shutdown signal received, initiating graceful shutdown...")

	// Cancel context to notify all components
	cancel()

	// Create shutdown context with timeout
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer shutdownCancel()

	// Shutdown API server first to stop accepting new requests
	if err := apiServer.Stop(shutdownCtx); err != nil {
		log.Error("Error stopping API server:", err)
	}

	// Stop queue manager to finish processing existing requests
	if err := queueManager.Stop(shutdownCtx); err != nil {
		log.Error("Error stopping queue manager:", err)
	}

	// Stop browser manager last
	if err := browserManager.Stop(shutdownCtx); err != nil {
		log.Error("Error stopping browser manager:", err)
	}

	log.Info("Graceful shutdown completed")
}
