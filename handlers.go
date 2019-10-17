package main

import (
	"log"
	"net/http"
)

//Health Check function
func Health(w http.ResponseWriter, r *http.Request) {
	log.Println("Health Endpoint")
	w.WriteHeader(200)
	w.Write([]byte("Health Check"))
}

//AddHandler Endpoint
func AddHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Add Handler Endpoint")
	w.WriteHeader(200)
	w.Write([]byte("Add Handler"))
}

//SubtractHandler Endpoint
func SubtractHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Subtract Handler Endpoint")
	w.WriteHeader(200)
	w.Write([]byte("Subtract Handler"))
}

//MultiplyHandler Endpoint
func MultiplyHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Multiply Handler Endpoint")
	w.WriteHeader(200)
	w.Write([]byte("Multiply Handler"))
}

//DivideHandler Endpoint
func DivideHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Division Handler")
	w.WriteHeader(200)
	w.Write([]byte("Divide Handler"))

}
