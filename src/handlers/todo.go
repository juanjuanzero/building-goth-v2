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
		th.Logger.Warn("error parsing date", "err", err)
		dueDate = time.Now()
	}

	var request todo.ToDoItem
	request.Task = todoAction
	request.Due = dueDate
	request.Id = fmt.Sprintf("%v,%v", todoAction, dueDate.String())

	todo.Add(request)
	th.Logger.Info(fmt.Sprintf("added: %+v", request))
	row := components.ToDoItemRow(request)
	row.Render(context.Background(), w)
}

func (th *TodoHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// read from the request
	id := r.PathValue("id")
	todo.Delete(id)
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte(`deleted`))
}
