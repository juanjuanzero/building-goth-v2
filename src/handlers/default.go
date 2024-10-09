package handlers

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/juanjuanzero/building-goth-v2/src/components"
)

type ServerHandler struct {
	Mux         *http.ServeMux
	Log         *slog.Logger
	TodoHandler *TodoHandler
}

func New(log *slog.Logger) *ServerHandler {
	srv := &ServerHandler{
		Mux: http.NewServeMux(),
		Log: log,
	}
	srv.bootstrap()
	return srv
}

func (sh *ServerHandler) bootstrap() {
	sh.TodoHandler = &TodoHandler{Logger: sh.Log, BaseRoute: "/todo"}
	sh.addRoutes()
}

// addRoutes requires services to be created and passed into it
func (sh *ServerHandler) addRoutes() {
	sh.Mux.HandleFunc("POST /todo/add", sh.TodoHandler.Add)
	sh.Mux.HandleFunc("GET /todo/{id}", sh.TodoHandler.Get)
	sh.Mux.HandleFunc("PUT /todo/{id}", sh.TodoHandler.Update)
	sh.Mux.HandleFunc("DELETE /todo/{id}", sh.TodoHandler.Delete)

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
