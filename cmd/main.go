package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/handlers"
	"github.com/vinnedev/http-server-go-boilerplate/internal/app"
	"github.com/vinnedev/http-server-go-boilerplate/internal/infrastructure/config"
	httpInterface "github.com/vinnedev/http-server-go-boilerplate/internal/interfaces/http"
)

func main() {
	// Initialize the Health service
	healthService := app.NewHealthService()

	// Initialize the HTTP handler
	healthHandler := httpInterface.NewHealthHandler(healthService)

	// Configure the mux for routing
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler.CheckHealth)

	// Cors
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"})

	// Initialize the HTTP server
	server := &http.Server{
		Addr: ":" + config.PORT,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(mux),
	}

	// Channel to capture system interrupt signals (SIGINT, SIGTERM)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		fmt.Printf("Server running on port %s ⚙️\n", config.PORT)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s: %v\n", config.PORT, err)
		}
	}()

	<-stop
	fmt.Println("\nShutting down gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	fmt.Println("Server exited")
}
