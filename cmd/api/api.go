package main

import (
	//"database/sql"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type application struct {
	config config // application has a config struct
}

type config struct {
	addr string // Address is stored inside config
}

func (app *application) mount() *chi.Mux {

	//Chi is a faster and more efficient router for handling request compared to http.NewServeMux()
	r := chi.NewRouter()

	// A good base middleware stack
	//Middleware is good for authentication and error handling of client requests
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	//No parenthesis is called function reference
	//It is used here because we are handling requests before server has started

	r.Route("/v1", func(r chi.Router) {
		//No parenthesis is called function reference
		//It is used here because we are handling requests before server has started
		r.Get("/health", app.healthCheckHandler)
		r.Get("/giga", app.gigaCheckHandler)
	})

	return r
}

func (app *application) run(mux *chi.Mux) error {

	srv := &http.Server{

		//Creates the server address and a http request handler
		Addr:    app.config.addr,
		Handler: mux,

		//Sets a time limit for responding to client and client reading the response
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	//log.Printf("Server has started, %s",
	//app.config.addr)

	return srv.ListenAndServe()
}
