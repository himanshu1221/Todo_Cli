package todo

import (
	"errors"
	"time"
)

// Declaring tthe structure of the Todos
type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

// Slice used to store dynamic data
type Todos []item

// Controller to add task in todo
func (t *Todos) Add(task string) {

	todo := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*t = append(*t, todo)
}

// Contorller to mark completed task
func (t *Todos) Completed(index int) error {

	ls := *t
	if index <= 0 || index > len(ls) {
		return errors.New("invalid text")
	}

	ls[index-1].CompletedAt = time.Now()
	ls[index-1].Done = true

	return nil
}

// Controller to delete a Todo
func (t *Todos) Delete(index int) error {

	ls := *t

	if index <= 0 || index > len(ls) {
		return errors.New("invalid text")
	}

	//Used slicing to remove the todo and connect it
	*t = append(ls[:index-1], ls[index:]...)

	return nil
}
