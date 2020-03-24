package router

import (
	"io"
	"net/http"
)

// Route : (unit -> string) -> (http.ResponseWrite, *http.Request -> unit)
func Route(logic func() string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, _ *http.Request) {
		response := logic()
		io.WriteString(w, response)
	}
}