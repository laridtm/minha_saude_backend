package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/laridtm/minha_saude/internal/env"
	"github.com/laridtm/minha_saude/internal/http"
	"github.com/laridtm/minha_saude/internal/repository/mongo"
)

func main() {
	log.Print("Starting API...")

	ctx, cancel := context.WithCancel(context.Background())

	mongoClient, err := mongo.NewClient(ctx, env.GetMongoURL())
	if err != nil {
		log.Printf("error on connect mongo: %s", err.Error())
	}
	defer func() {
		if err := mongoClient.Disconnect(ctx); err != nil {
			log.Printf("error on disconnect mongo: %s", err.Error())
		}
	}()

	handler := http.NewHandler()
	server := http.NewServer("8080", handler)
	server.Start()

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)
	<-stopChan

	cancel()
	server.ShutDown()
}
