package main

import (
	"log"
	"net/http"
)

func main() {
	a := Initialize()
	log.Println("Starting Router")
	http.ListenAndServe(a.port, a.router)
}
