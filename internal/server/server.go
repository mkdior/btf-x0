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
		Addr:    "172.186.0.2:8082",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("failed starting server...\n%s", err)
	}

	mux.HandleFunc("POST /user/create", handleUserCreate)
	mux.HandleFunc("POST /merkle/build", handleMerkleBuild)
	mux.HandleFunc("GET /merkle/proof/generate", handleMerkleProofGenerate)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT)
	<-sig
	log.Print("Shutting down server... gn")

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err = server.Shutdown(ctx)
	if err != nil {
		log.Fatalf("failed shutting down server...\n%s", err)
	}
}
