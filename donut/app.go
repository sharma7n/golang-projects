package main

import (
	"log"
	"net/http"
	"os"

	"database/sql"
	_ "github.com/lib/pq"

	"app/lib/server"
	"app/gen/donut"
	"app/src/effect"
	"app/src/store"
)

func index() server.Route {
	return func() *server.Reply {
		return server.Text("Hello, world!")
	}
}

func listRoute(getDonutList effect.GetDonutList) server.Route {
	return func() *server.Reply {
		list, err := getDonutList.GetDonutList()
		if err != nil {
			return server.Error(err)
		}
		return server.Proto(&list)
	}
}

func addRoute(addDonut effect.AddDonut) server.Route {
	return func() *server.Reply {
		err := addDonut.AddDonut(donut.Donut{Shape: donut.Shape_RING})
		if err != nil {
			return server.Error(err)
		}
		return server.Text("Added ring donut")
	}
}

func configureHTTPHandlers(db *sql.DB) {
	store := store.Store{DB: db}
	http.HandleFunc("/", server.ToHandler(index()))
	http.HandleFunc("/list", server.ToHandler(listRoute(store)))
	http.HandleFunc("/add", server.ToHandler(addRoute(store)))
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