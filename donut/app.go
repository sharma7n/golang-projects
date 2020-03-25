package main

import (
	"database/sql"
	_ "github.com/lib/pq"

	"encoding/json"
	"log"
	"net/http"
	"os"
)

// Shape ...
type Shape int

const (
	// Ring ...
	Ring Shape = 0

	// Hole ...
	Hole Shape = 1
)

// Donut ...
type Donut struct {
	Shape Shape
}

func configureHTTPHandlers(db *sql.DB) {
	http.HandleFunc("/", route(func() ([]byte, error) {
		return []byte("Hello, world!"), nil
	}))

	http.HandleFunc("/list", route(func() ([]byte, error) {
		donuts, err := getDonuts(db)
		if err != nil {
			return nil, err
		}

		return json.Marshal(donuts)
	}))
}

func main() {
	// Program arguments
	address := os.Args[1] // full host & port (e.g. "0.0.0.0:3000")
	connStr := os.Args[2] // postgres connection string

	// Configure the database connection using the given connection string
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Build the application using the database connection
	configureHTTPHandlers(db)

	// Serve the application using the given address
	err = http.ListenAndServe(address, nil)
	log.Fatal(err)
}

func getDonuts(db *sql.DB) (donuts []Donut, err error) {
	rows, err := db.Query("SELECT * FROM Donut;")
	if err != nil {
		return
	}
	defer rows.Close()

	
	for rows.Next() {
		var shape Shape
		if err = rows.Scan(&shape); err != nil {
			return
		}
		donut := Donut{Shape: shape}
		donuts = append(donuts, donut)
	}
	
	if err = rows.Err(); err != nil {
		return
	}

	return
}

func route(logic func() ([]byte, error)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL)

		response, err := logic()
		if err != nil {
			log.Fatal(err)
		}

		w.Write(response)
	}
}