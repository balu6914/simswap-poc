package main

import (
	"simswap-poc/config"
	"simswap-poc/server"
)

func main() {
	// Load config
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		panic(err)
	}

	// Start server
	server.StartServer(cfg)
}
