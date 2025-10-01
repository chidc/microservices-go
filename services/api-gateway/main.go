package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"ride-sharing/shared/env"
)

var (
	httpAddr = env.GetString("HTTP_ADDR", ":8081")
)

func main() {
	log.Println("Starting API Gateway")

	mux := http.NewServeMux()

	mux.HandleFunc("POST /trip/preview", enableCORS(handleTripPreview))
	mux.HandleFunc("/ws/drivers", handleDriversWebSocket)
	mux.HandleFunc("/ws/riders", handleRidersWebSocket)
	server := &http.Server{
		Addr:    httpAddr,
		Handler: mux,
	}

	certFile := "ride-sharing/services/api-gateway/ca.crt" // TODO: set your cert path
	keyFile := "ride-sharing/services/api-gateway/ca.key"

	serverErrors := make(chan error, 1)
	go func() {
		log.Printf("server listening on %s", httpAddr)
		serverErrors <- server.ListenAndServeTLS(certFile, keyFile)
	}()
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		log.Fatalf("server error: %v", err)
	case sig := <-shutdown:
		log.Printf("shutdown signal received: %v", sig)
		if err := server.Close(); err != nil {
			log.Printf("could not close server: %v", err)
		}
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			log.Printf("could not gracefully shutdown the server: %v", err)
			server.Close()
		}
	}
}
