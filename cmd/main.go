package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/markable/internal/dependency"
	"github.com/markable/internal/http/swagger"
	"github.com/markable/internal/pkg/config"
)

func main() {
	cfgOpt := getConfigOptions()

	cfg, err := dependency.NewConfig(cfgOpt) // Pass Options directly
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	// Setup database connection
	db, err := dependency.NewDatabaseConfig(cfg)
	if err != nil {
		log.Fatalf("failed to create connection for database: %v", err)
	}
	err = db.Ping(context.Background())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize the dependencies
	api, err := dependency.NewMarkableApi(cfg, db)
	if err != nil {
		log.Fatalf("failed to create dependencies: %v", err)
	}

	e := echo.New()
	e.HideBanner = true

	api.SetupMiddleware(e)
	swagger.SetupSwagger(cfg, e)
	api.SetupRoutes(e)

	// Start server in a goroutine
	go func() {
		e.Logger.Info(e.Start(fmt.Sprintf("0.0.0.0:%d", cfg.AppPort)))
	}()

	// Graceful shutdown handling
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Server is shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
	log.Println("Server gracefully stopped")
}

func getConfigOptions() config.Options {
	cfgSource := os.Getenv(config.SourceKey)
	if cfgSource == "" {
		cfgSource = config.SourceEnv
	}
	return config.Options{
		ConfigFileSource: cfgSource,
		ConfigFile:       ".env",
	}
}
