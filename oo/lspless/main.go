// We cannot replace (or even assign) a variable of a supertype, e.g. "Task"
// with a subtype, e.g. "Extract".
package main

import (
	"fmt"
)

type Task struct {
	Name string
}

type Extract struct {
	Task
	Source string
}

func (e *Extract) String() string {
	return fmt.Sprintln("[%s] extract", e.Name)
}

type Transform struct {
	Task
}

func (t *Transform) String() string {
	return fmt.Sprintln("[%s] transform", t.Name)
}

func main() {
	// ./main.go:32:6: cannot use Extract{} (type Extract) as type Task in assignment
	var task Task = Extract{}

	// # command-line-arguments
	// ./main.go:31:10: cannot use Extract{...} (type Extract) as type Task in slice literal
	// ./main.go:31:11: cannot use promoted field Task.Name in struct literal of type Extract
	// ./main.go:32:12: cannot use Transform{...} (type Transform) as type Task in slice literal
	// ./main.go:32:13: cannot use promoted field Task.Name in struct literal of type Transform
	tasks := []Task{
		Extract{Name: "E"},
		Transform{Name: "T"},
	}
}
