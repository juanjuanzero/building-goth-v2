package main

import (
	"fmt"
	"log/slog"
	"net"
	"net/http"
)

type server struct {
	http *http.Server
}

func NewServer(host string, port string, handler http.Handler) *server {
	var httpServer server
	httpServer.http = &http.Server{
		Addr:    net.JoinHostPort(host, port),
		Handler: handler,
	}
	return &httpServer
}

// addRoutes requires services to be created and passed into it
func addRoutes(
	mux *http.ServeMux,
) {
	mux.HandleFunc("/", HandleHome)
	slog.Debug("added all routes")
}

// TODO move to handlers
// reply with the home page
func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello There</h1>")
}

func main() {
	// create handler
	handle := http.NewServeMux()
	// add routes
	addRoutes(handle)
	// start the server
	slog.Debug("new server")
	svr := NewServer("localhost", "8080", handle)
	svr.http.ListenAndServe()
}
