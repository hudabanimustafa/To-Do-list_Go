package main

// Todo represents a task with an ID, title, and completion status.
type Todo struct {
	ID        int
	Title     string
	Completed bool
}

// todos stores the list of all tasks.
var todos []Todo

// nextID is used to assign a unique ID to each new task.
var nextID int = 1

// AddTodo adds a new task to the todo list.
func AddTodo(title string) {
	todo := Todo{
		ID:        nextID,
		Title:     title,
		Completed: false,
	}
	todos = append(todos, todo)
	nextID++
}

// CompleteTodo marks a task as completed based on its ID.
func CompleteTodo(id int) {
	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Completed = true
			break
		}
	}
}

// DeleteTodo removes a task from the todo list based on its ID.
func DeleteTodo(id int) {
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			break
		}
	}
}

// GetTodos returns the list of all tasks.
func GetTodos() []Todo {
	return todos
}
