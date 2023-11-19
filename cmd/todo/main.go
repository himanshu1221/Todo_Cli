package main

import (
	"flag"
	"fmt"
	"os"

	todo "github.com/himanshu1221/Todo_Cli"
)

const (
	todoFile = ".todo.json"
)

func main() {

	add := flag.Bool("add", false, "add a new todo")

	flag.Parse()

	//pointing to package todo that we have created ie todo.go
	todos := &todo.Todos{}

	if err := todos.Load(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case *add:
		todos.Add("Sample Todo")
		err := todos.Store(todoFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stdout, "Invalid Command")
		os.Exit(0)
	}
}
