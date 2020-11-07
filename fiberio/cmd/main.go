package main

import (
	server "fiberio/internal/app/fiberio"
)

func main() {
	cfg := server.NewConfig()
	server.Start(cfg)
}
