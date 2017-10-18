package datasources

import (
	"encoding/json"
	"io"
	"log"
)

// writeJSON writes our structure out in a nice way.
func writeJSON(i interface{}, w io.Writer) {
	b, _ := json.MarshalIndent(i, "", "  ")
	b = append(b, []byte("\n")...)
	_, err := w.Write(b)
	if err != nil {
		log.Fatal(err)
	}
}
