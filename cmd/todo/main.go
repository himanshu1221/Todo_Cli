package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	todo "github.com/himanshu1221/Todo_Cli"
)

const (
	todoFile = ".todo.json"
)

func main() {

	add := flag.Bool("add", false, "Add a new todo")
	complete := flag.Int("complete", 0, "Mark a todo as completed")
	del := flag.Int("del", 0, "Delete a todo")
	list := flag.Bool("list", false, "List all the todos")
	helpFlag := flag.Bool("help", false, "Display help message")
	flag.Parse()

	//pointing to package todo that we have created ie todo.go
	todos := &todo.Todos{}

	if err := todos.Load(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	switch {
	case *add:
		task, err := getInput(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		todos.Add(task)
		err = todos.Store(todoFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}
	case *complete > 0:
		err := todos.Completed(*complete)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		err = todos.Store(todoFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	case *del > 0:
		err := todos.Delete(*del)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		err = todos.Store(todoFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	case *list:
		todos.List()
	case *helpFlag:
		displayHelp()
	default:
		fmt.Fprintln(os.Stdout, "Invalid Command use -help for usage")
		os.Exit(0)
	}
}

//Get input from the user

func getInput(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}
	if len(scanner.Text()) == 0 {
		return "", errors.New("empty todo")
	}
	return scanner.Text(), nil
}

// Display Help Message
func displayHelp() {
	fmt.Println("\nAvailable Commands:")
	fmt.Println("  -add `Your text here` - to add todo")
	fmt.Println("  -complete `Number of todo` - to mark a todo as completed")
	fmt.Println("  -del `Number of todo` - To delete a todo")
	fmt.Println("  -list - To list add the todos")

	// Exit the program
	os.Exit(0)
}
