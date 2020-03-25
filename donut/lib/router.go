package lib

import (
	"net/http"
)

// Route : (unit -> []byte]) -> (http.ResponseWrite, *http.Request -> unit)
func Route(logic func() []byte) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, _ *http.Request) {
		response := logic()
		w.Write(response)
	}
}