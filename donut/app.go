package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"../donut/lib"
	"../donut/src"
)

func configureHTTPHandlers() {
	http.HandleFunc("/", lib.Route(func() []byte {
		return []byte("Hello, world!")
	}))

	http.HandleFunc("/list", lib.Route(func() []byte {
		donuts := []src.Donut{
			src.Donut{
				Shape: src.Ring,
			},
			src.Donut{
				Shape: src.Hole,
			},
		}

		response, err := json.Marshal(donuts)
		if err != nil {
			panic(err)
		}

		return response
	}))
}

func main() {
	port := os.Args[1]
	fmt.Println("Port:", port)
	
	// Serve on localhost
	address := fmt.Sprintf(":%s", port)

	configureHTTPHandlers()

	// Serve application
	err := http.ListenAndServe(address, nil)
	log.Fatal(err)
}