package main

import (
	"log"
	"net/http"
	"os"

	"database/sql"
	_ "github.com/lib/pq"

	"github.com/golang/protobuf/proto"

	"app/gen/donut"
)

func configureHTTPHandlers(db *sql.DB) {
	http.HandleFunc("/", route(func() ([]byte, error) {
		return []byte("Hello, world!"), nil
	}))

	http.HandleFunc("/list", route(func() ([]byte, error) {
		list, err := getDonutList(db)
		if err != nil {
			return nil, err
		}

		return proto.Marshal(&list)
	}))

	http.HandleFunc("/add", route(func() ([]byte, error) {
		err := addDonut(db)
		if err != nil {
			return nil, err
		}

		return []byte("Added donut"), nil
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

func getDonutList(db *sql.DB) (list donut.DonutList, err error) {
	rows, err := db.Query("SELECT * FROM Donut;")
	defer rows.Close()
	if err != nil {
		return
	}
	
	for rows.Next() {
		var shape donut.Shape
		if err = rows.Scan(&shape); err != nil {
			return
		}
		donut := donut.Donut{Shape: shape}
		list.Donuts = append(list.Donuts, &donut)
	}
	
	if err = rows.Err(); err != nil {
		return
	}

	return
}

func addDonut(db *sql.DB) error {
	rows, err := db.Query("INSERT INTO Donut VALUES (1)")
	defer rows.Close()
	return err
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