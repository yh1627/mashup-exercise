package mashupserver

import (
	"log"
	"net/http"

	"encoding/json"
	"git.playchat.net/playchat/mashup-exercise/simpleserver"
)

type R struct {
	Value int
}

type Result struct {
	Value int `json:"final_result"`
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

	var url_list = [2]string{"http://127.0.0.1:10000/value", "http://127.0.0.1:11000/value"}

	sum := 0
	status := 200
	for i := 0; i < len(url_list); i++ {

		req, err := http.NewRequest("GET", url_list[i], nil)

		if err != nil {
			log.Fatal("Error making request", err)
			status = 503
		}

		client := &http.Client{}
		resp, err := client.Do(req)

		if err != nil {
			log.Fatal("Error geting response", err)
			status = 503
		}

		defer resp.Body.Close()

		var res R
		if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
			log.Fatalf("Error decoding the resule", err)
			status = 503
		}

		sum = sum + int(res.Value)
	}

	result := Result{Value: sum}

	finale_result, _ := json.Marshal(result)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	w.Write([]byte(finale_result))
}
