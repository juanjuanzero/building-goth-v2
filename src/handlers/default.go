package handlers

import (
	"fmt"
	"log/slog"
	"net/http"
)

type ServerHandler struct {
	Mux *http.ServeMux
	Log *slog.Logger
}

func New(log *slog.Logger) *ServerHandler {
	return &ServerHandler{
		Mux: http.NewServeMux(),
		Log: log,
	}
}

// addRoutes requires services to be created and passed into it
func (sh *ServerHandler) AddRoutes() {
	sh.Mux.HandleFunc("/", HandleHome)
	sh.Log.Info("added all routes")
}

func (sh *ServerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sh.Mux.ServeHTTP(w, r)
}

// TODO move to handlers
// reply with the home page
func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello There</h1>")
}
