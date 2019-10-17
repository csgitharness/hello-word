package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Health Check function
func Health(w http.ResponseWriter, r *http.Request) {
	log.Println("Health Endpoint")
	w.WriteHeader(200)
	w.Write([]byte("Health Check"))
}

//NewRouter Function
func (a *App) NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/health", Health).Methods("GET")
	return r
}
