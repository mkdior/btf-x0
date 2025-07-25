package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Start() {
	mux := http.NewServeMux()
	server := http.Server{
		Addr:    "127.0.0.1:8082",
		Handler: mux,
	}
	log.Print("Starting server...")
	go func() {
		if err := server.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {
			log.Fatalf("failed starting server...\n%s", err)
		}
	}()

	mux.HandleFunc("POST /user/create", handleUserCreate)
	mux.HandleFunc("POST /merkle/build", handleMerkleBuild)
	mux.HandleFunc("GET /merkle/proof/generate", handleMerkleProofGenerate)

	sig := make(chan os.Signal, 2)
	signal.Notify(sig, os.Interrupt, syscall.SIGINT)
	<-sig
	log.Print("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
		log.Fatalf("failed shutting down server...\n%s", err)
	}
}
