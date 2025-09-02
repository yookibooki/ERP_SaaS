package main

import (
	"log"

	"github.com/yookibooki/ERP_SaaS/internal/app/inventory"
	"github.com/yookibooki/ERP_SaaS/internal/pkg/config"
)

func main() {
	// Load configuration
	cfg, err := config.Load("inventory")
	if err != nil {
		log.Fatalf("could not load configuration: %v", err)
	}

	// Instantiate the server
	server := inventory.NewServer(cfg)

	// Start the server
	log.Println("Inventory Service starting...")
	if err := server.Run(); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
