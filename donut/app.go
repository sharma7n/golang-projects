package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"../donut/src"
)

func configureHTTPHandlers() {
	http.HandleFunc("/", router.Route(func() string {
		return "Hello, world!"
	}))
}

func main() {
	// Default port to 3000
	port, success := os.LookupEnv("APP_PORT")
	if !success {
		port = "3000"
	}
	
	// Serve on localhost
	address := fmt.Sprintf(":%s", port)

	configureHTTPHandlers()

	// Serve application
	err := http.ListenAndServe(address, nil)
	log.Fatal(err)
}