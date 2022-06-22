package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/JanCalebManzano/go-microservices/services/user/http/servers"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	s := servers.NewGinServer(os.Getenv("HTTP_PORT"))

	go func() {
		if err := s.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt)
	signal.Notify(sigCh, os.Kill)
	sig := <-sigCh
	fmt.Println("Graceful shutdown", sig)

	s.Shutdown()
}
