package main

import (
	"log"

	"github.com/gigak23/ecom/internal/env"
)

func main() {

	cfg := config{
		addr: env.GetString("ADDR", ":8080"), // Create config with addr
	}
	app := &application{
		config: cfg, // Pass config into application struct
	}

	//We create the setup for handling reqeust from client
	mux := app.mount()
	//Then we start the server and continously look for client request
	log.Fatal(app.run(mux))
}
