package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/cendaar/fizzbuzz/api"
)

func main() {
	server := api.NewServer()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		err := server.Start()
		if err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	<-quit
}
