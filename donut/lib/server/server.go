package server


import (
	"net/http"
	"log"

	"github.com/golang/protobuf/proto"
)

// Route represents a server function that may perform some number of side effects and will return a *Reply.
type Route func() *Reply

// Reply represents the return value of an HTTP route in bytes.
type Reply struct {
	Bytes []byte
	Error error
}

// Error constructs a *Reply from a given error.
func Error(err error) *Reply {
	return &Reply{Bytes: nil, Error: err}
}

// Text constructs a *Reply from a given string.
func Text(text string) *Reply {
	return &Reply{Bytes: []byte(text), Error: nil}
}

// Proto constructs a *Reply from a protobuf Message.
func Proto(pb proto.Message) *Reply {
	bytes, err := proto.Marshal(pb)
	return &Reply{Bytes: bytes, Error: err}
}

// ToHandler promotes a Route function to an http HandleFunc.
func ToHandler(logic Route) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL)

		reply := logic()
		if reply.Error != nil {
			log.Fatal(reply.Error)
		}

		w.Write(reply.Bytes)
	}
}