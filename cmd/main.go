package main

import (
	"time"
)

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

	api.Run()
}
