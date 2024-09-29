package main

import (
	"log/slog"
	"net"
	"net/http"
	"os"

	"github.com/juanjuanzero/building-goth-v2/src/handlers"
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

func main() {
	// create handler
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	handler := handlers.New(logger)
	// start the server
	logger.Info("new server")
	svr := NewServer("localhost", "8080", handler)
	svr.http.ListenAndServe()
}
