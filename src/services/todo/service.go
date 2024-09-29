package todo

import (
	"fmt"
	"time"
)

type NoTodoItem struct {
	ExpectedToDoItem string
}

func (nti *NoTodoItem) Error() string {
	return fmt.Sprintf("no to do item with id %v", nti.ExpectedToDoItem)
}

var ToDoItems map[string]ToDoItem = make(map[string]ToDoItem, 0)

type ToDoItem struct {
	Id   string    `json:"id"`
	Task string    `json:"task"`
	Due  time.Time `json:"due"`
}

func Add(todo ToDoItem) {
	ToDoItems[todo.Id] = todo
}

func Get(id string) (ToDoItem, error) {
	retrieved, ok := ToDoItems[id]
	if !ok {
		return ToDoItem{}, &NoTodoItem{ExpectedToDoItem: id}
	}
	return retrieved, nil
}

func Update(todo ToDoItem) (ToDoItem, error) {
	retrieved, err := Get(todo.Id)
	if err != nil {
		return ToDoItem{}, err
	}

	// update, not necessary but will be later
	retrieved.Due = todo.Due
	retrieved.Task = todo.Task

	// save it
	ToDoItems[todo.Id] = retrieved
	return retrieved, nil
}
