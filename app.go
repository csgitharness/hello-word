package main

import (
	"log"

	"github.com/gorilla/mux"
)

//App - Application config
type App struct {
	router *mux.Router
	port   string
}

//Initialize App
func Initialize() App {
	log.Println("Initializing the Application")
	app := App{}
	app.port = ":9000"
	app.router = app.NewRouter()

	return app
}

//NewRouter Function
func (a *App) NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/health", Health).Methods("GET")
	r.HandleFunc("/add", AddHandler).Methods("POST")
	return r
}
