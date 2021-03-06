package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	todo "github.com/VictorMilhomem/todoCli/src/entity"
	"io"
	"os"
	"strings"
)

const (
	todoFile = ".todos.json"
)

func main() {

	add := flag.Bool("add", false, "add a new todo")
	complete := flag.Int("complete", 0, "mark a todo as completed")
	del := flag.Int("delete", 0, "delete a todo")
	list := flag.Bool("list", false, "show all todos")
	flag.Parse()

	todos := &todo.Todos{}

	if err := todos.Load(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case *add:
		task, err := getInput(os.Stdin, flag.Args()...)
		checkErr(err)
		todos.Add(task)
		err = todos.Store(todoFile)
		checkErr(err)

	case *complete > 0:
		err := todos.Complete(*complete)
		checkErr(err)
		err = todos.Store(todoFile)
		checkErr(err)
	case *del > 0:
		err := todos.Delete(*del)
		checkErr(err)
		err = todos.Store(todoFile)
		checkErr(err)
	case *list:
		todos.Print()
	default:
		fmt.Fprintln(os.Stdout, "invalid command")
		os.Exit(0)

	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func getInput(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	scanner := bufio.NewScanner(r)
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		return "", err
	}

	text := scanner.Text()

	if len(text) == 0 {
		return "", errors.New("empty todo is not permited")
	}

	return text, nil
}
