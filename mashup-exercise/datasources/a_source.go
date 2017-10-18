package datasources

import (
	"log"
	"net/http"

	"git.playchat.net/playchat/mashup-exercise/simpleserver"
)

// A is the return value from a SourceA
type A struct {
	Value int
}

type SourceA struct {
	*simpleserver.Server
}

// NewSourceA creates a new Source "A", listening on port 10000.
func NewSourceA() *simpleserver.Server {
	srv, err := simpleserver.ServeAddr(":10000", []*simpleserver.PathHandler{
		{Path: "/value", Handler: func(w http.ResponseWriter, r *http.Request) {
			writeJSON(&A{Value: 45}, w)
		}},
	})
	if err != nil {
		log.Fatal(err)
	}
	return srv
}
