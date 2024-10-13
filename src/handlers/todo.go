package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/juanjuanzero/building-goth-v2/src/components"
	"github.com/juanjuanzero/building-goth-v2/src/services/todo"
)

type TodoRequest struct {
	Todo todo.ToDoItem `json:"todoItem"`
}

type TodoHandler struct {
	Logger    *slog.Logger
	BaseRoute string
}

// no obsolete
func (th *TodoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	th.Logger.Info(fmt.Sprintf("uri: %v, method: %v\n", r.RequestURI, r.Method))
	if r.RequestURI == fmt.Sprintf("/%v/add", th.BaseRoute) {
		th.Add(w, r)
	}

	if r.Method == "GET" {
		th.Get(w, r)
	}

	if r.Method == "POST" {
		th.Update(w, r)
	}
}

func (th *TodoHandler) Update(w http.ResponseWriter, r *http.Request) {
	// unmarshall
	var item todo.ToDoItem
	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		th.Logger.Info("err")
	}
	json.Unmarshal(data, &item)

	updatedItem, err := todo.Update(item)
	if err != nil {
		th.Logger.Error(err.Error())
	}
	itemB, err := json.Marshal(updatedItem)
	if err != nil {
		th.Logger.Error(err.Error())
	}
	w.Write(itemB)
}

func (th *TodoHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	data := todo.GetAll()
	table := components.ToDoTable(data)
	page := components.Layout(table)
	page.Render(context.Background(), w)
}

func (th *TodoHandler) Get(w http.ResponseWriter, r *http.Request) {
	// read the request URI to retrieve the id
	id := r.PathValue("id")
	item, err := todo.Get(id)
	if err != nil {
		th.Logger.Error(err.Error())
		w.WriteHeader(http.StatusNotFound)
	}

	itemB, err := json.Marshal(item)
	if err != nil {
		th.Logger.Error(err.Error())
	}
	w.Write(itemB)

}

func (th *TodoHandler) Add(w http.ResponseWriter, r *http.Request) {
	// parse the request body to get the item
	if err := r.ParseForm(); err != nil {
		th.Logger.Info("err")
	}

	todoAction := r.FormValue("task")
	due := r.FormValue("due")
	dueDate, err := time.Parse(time.DateOnly, due)
	if err != nil {
		th.Logger.Error("err")
	}

	var request todo.ToDoItem
	request.Task = todoAction
	request.Due = dueDate

	todo.Add(request)
	th.Logger.Info(fmt.Sprintf("added: %+v", request))
}

func (th *TodoHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// read from the request
	id := r.PathValue("id")
	todo.Delete(id)
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte(`deleted`))
}
