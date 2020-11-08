package main

import (
	server "fiberio/internal/app/fiberio"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
)

func main() {
	cfg, err := server.NewConfig()
	if err != nil {
		log.Fatalf("Error loading conf: %v", err)
		os.Exit(-1)
	}
	if err := server.Start(cfg); err != nil {
		log.Fatalf("Error starting server: %v", err)
		os.Exit(-1)
	}
}
