package handlers

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/juanjuanzero/building-goth-v2/src/components"
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
	sh.Mux.Handle("/static/*", http.StripPrefix("/static", HandleStatic()))
	sh.Mux.HandleFunc("/", HandleHome)
	sh.Log.Info("added all routes")
}

func (sh *ServerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sh.Mux.ServeHTTP(w, r)
}

// TODO move to handlers
// reply with the home page
func HandleHome(w http.ResponseWriter, r *http.Request) {
	home := components.Layout()
	home.Render(context.TODO(), w)
}
