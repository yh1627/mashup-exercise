package simpleserver

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	testBody = "Test Body"
)

func Test_Simple_Get(t *testing.T) {
	srv, err := Serve([]*PathHandler{
		&PathHandler{"/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(testBody))
		}},
	})
	require.NoError(t, err)

	defer srv.MustShutdown(context.Background())

	c := &http.Client{}

	res, err := c.Get(srv.URL)
	require.NoError(t, err)

	body, err := ioutil.ReadAll(res.Body)
	require.NoError(t, err)
	require.Equal(t, testBody, string(body))
}

func Test_Bad_Address(t *testing.T) {
	_, err := ServeAddr(":-2020", []*PathHandler{
		&PathHandler{"/", func(w http.ResponseWriter, r *http.Request) {
			log.Print("HERE")
		}},
	})
	require.Error(t, err)
}
