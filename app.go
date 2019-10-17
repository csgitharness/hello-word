package main

import "github.com/gorilla/mux"

//App - Application config
type App struct {
	router *mux.Router
	port   string
}
