package main

import "net/http"

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to server my"))
}

func (app *application) gigaCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to giga page no diddy"))
}
