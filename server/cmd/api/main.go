package main 

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/TutorLink/server/internal/app"
	"github.com/TutorLink/server/pkg/config"
	"github.com/TutorLink/server/pkg/db"
	"github.com/TutorLink/server/pkg/logger"
)

func main(){
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize logger
	l := logger.New(cfg.Environment)
	defer func() {
		if err := l.Sync(); err != nil {
			log.Printf("Failed to sync logger: %v", err)
		}
	}()

	// Connect to the primary database
	database, err := db.New(cfg.Database)
	if err != nil {
		l.Fatal("Failed to connect to database", err)
	}

	// Connect to Redis
	redis, err := db.NewRedis(cfg.Redis)
	if err != nil {
		l.Fatal("Failed to connect to Redis", err)
	}

	// Initialize the application with dependencies
	application := app.New(cfg, l, database, redis)

	// Start the server in a separate goroutine
	go func(){
		if err := application.Start(); err != nil {
			l.Fatal("Failed to start server", err)
		}
	}()

	// Set up channel to listen for interrupt or terminate signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Block until a signal is received
	<-quit

	// Create a context with timeout for the shutdown process
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt to gracefully shutdown the application
	if err := application.Shutdown(ctx); err != nil {
		l.Error("Server forced to shutdown", err)
	}

	l.Info("Server exiting")
}
