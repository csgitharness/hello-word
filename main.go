package main

import (
	"log"
	"net/http"
)

//Initialize
func Initialize() App {
	log.Println("Initializing the Application")
	app := App{}
	app.port = ":9000"
	app.router = app.NewRouter()

	return app
}

func main() {
	a := Initialize()
	log.Println("Starting Router")
	http.ListenAndServe(a.port, a.router)
}
