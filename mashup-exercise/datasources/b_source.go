package datasources

import (
	"log"
	"net/http"

	"git.playchat.net/playchat/mashup-exercise/simpleserver"
)

// B is the return value for a source B.
type B struct {
	Value int
}

// SourceB is a http service that returns a JSON chunk like: {"IntValue": "3030"}.  Note that the type of the
// value returned is an integer, even though it is stored as a string.
type SourceB struct {
	*simpleserver.Server
}

// NewSourceB creates a new SourceB
func NewSourceB() *simpleserver.Server {
	srv, err := simpleserver.ServeAddr(":11000", []*simpleserver.PathHandler{
		{Path: "/value", Handler: func(w http.ResponseWriter, r *http.Request) {
			writeJSON(&B{Value: 466}, w)
		}},
	})
	if err != nil {
		log.Fatal(err)
	}
	return srv
}
