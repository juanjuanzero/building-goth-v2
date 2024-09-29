package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/juanjuanzero/building-goth-v2/src/services/todo"
)

type TodoRequest struct {
	Todo todo.ToDoItem `json:"todoItem"`
}

type TodoHandler struct {
	Logger    *slog.Logger
	BaseRoute string
}

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

func (th *TodoHandler) Get(w http.ResponseWriter, r *http.Request) {
	// read the request URI to retrieve the id
	id := r.PathValue("id")
	item, err := todo.Get(id)
	if err != nil {
		th.Logger.Error(err.Error())
	}

	itemB, err := json.Marshal(item)
	if err != nil {
		th.Logger.Error(err.Error())
	}
	w.Write(itemB)

}

func (th *TodoHandler) Add(w http.ResponseWriter, r *http.Request) {
	// parse the request body to get the item
	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		th.Logger.Info("err")
	}
	var request TodoRequest
	err = json.Unmarshal(data, &request)

	if err != nil {
		th.Logger.Info("err")
	}
	todo.Add(request.Todo)
	th.Logger.Info(fmt.Sprintf("added: %+v", request.Todo))
}
