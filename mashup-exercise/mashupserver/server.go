package mashupserver

import (
	"log"
	"net/http"

	"fmt"
	"git.playchat.net/playchat/mashup-exercise/simpleserver"
	"encoding/json"
)

type R struct {
	Value int
}

// MashupServer is a simple server that combines 2 data sources together.
type MashupServer struct {
	*simpleserver.Server
}

// New creates a new mashup server that combines Source A and Source B.
func New() *MashupServer {
	srv, err := simpleserver.ServeAddr(":12000", []*simpleserver.PathHandler{
		{Path: "/result", Handler: handle}})
	if err != nil {
		log.Fatal(err)
	}
	return &MashupServer{srv}
}

// Does the actual work to combine the two data sources:  This is hard-coded, just to show as an example.
func handle(w http.ResponseWriter, r *http.Request) {
	url1 := "http://127.0.0.1:10000/value"
	url2 := "http://127.0.0.1:11000/value"

	req1, err1 := http.NewRequest("GET", url1, nil)
	req2, err2 := http.NewRequest("GET", url2, nil)

	if err1 != nil || err2 != nil {
		log.Fatal("Error making request: ", err1, err2)
		return
	}

	client := &http.Client{}
	resp1, err := client.Do(req1)
	resp2, err := client.Do(req2)
	if err != nil {
		log.Fatal("Error getting response: ", err)
		return
	}

	defer resp1.Body.Close()
	defer resp2.Body.Close()

	var res1 R
	var res2 R

	if err := json.NewDecoder(resp1.Body).Decode(&res1); err != nil {
		log.Println(err)
	}
	if err := json.NewDecoder(resp2.Body).Decode(&res2); err != nil {
		log.Println(err)
	}

	result := fmt.Sprintf("{\n   \"final_result\":%d\n}\n", int(res1.Value) + int(res2.Value))
	w.Write([]byte(result))
}
