package main

import (
	"log"
	"time"
)

// TODO: Create a controller thing that auto assign
// Handlers to function as app controllers
func main() {
	// WebApplication(port, ctx)
	// Handles the below code;
	config := config{
		addr:     ":8080",
		dbConfig: dbConfig{},
	}

	api := application{
		config,
		time.Time{},
	}

	if err := api.run(); err != nil {
		// Exit with Fx
		log.Printf("Server has failed: %s", err.Error())
	}

}
