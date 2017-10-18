package simpleserver

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const (
	timeout = 15 * time.Second
)

// Server is a very simple HTTP server.
type Server struct {
	*http.Server
	Addr *net.TCPAddr
	URL  string
}

// PathHandler is a path + function for that path.
type PathHandler struct {
	Path    string
	Handler func(w http.ResponseWriter, r *http.Request)
}

// Serve creates an HTTP server for the specific port.
func Serve(h []*PathHandler) (*Server, error) {
	return ServeAddr("127.0.0.1:0", h)
}

// ServeAddr serves a specific address like ":8080" with the path handlers supplied.
func ServeAddr(addr string, h []*PathHandler) (*Server, error) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	r := mux.NewRouter()
	for _, ph := range h {
		r.HandleFunc(ph.Path, ph.Handler)
	}
	srv := &http.Server{
		Handler:      r,
		WriteTimeout: timeout,
		ReadTimeout:  timeout,
	}

	go func() {
		if err := srv.Serve(listener); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	URL := "http://" + listener.Addr().String()
	rsrv := &Server{srv, listener.Addr().(*net.TCPAddr), URL}

	return rsrv, nil
}

// MustShutdown shuts down the server, or panics.
func (ss *Server) MustShutdown(ctx context.Context) {
	if err := ss.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
