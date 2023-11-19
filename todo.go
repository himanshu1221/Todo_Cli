package todo

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

// Declaring the structure of the Todos
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

// Load the file ie load  the todos
func (t *Todos) Load(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if len(file) == 0 {
		return err
	}
	err = json.Unmarshal(file, t)
	if err != nil {
		return nil
	}
	return nil
}

// Store a todo
func (t *Todos) Store(filename string) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}
