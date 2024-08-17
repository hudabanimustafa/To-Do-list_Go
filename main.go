package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

// Todo represents a task with an ID, title, and completion status.
type Todo struct {
	ID        int
	Title     string
	Completed bool
}

// Slice to store the list of all tasks.
var todos []Todo

// Variable to keep track of the next unique ID for each new task.
var nextID int = 1

func main() {
	// Serve static files (CSS, JS, images) from the "static" directory.
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Define route handlers for various endpoints.
	http.HandleFunc("/", indexHandler)         // Root route to display the To-Do list.
	http.HandleFunc("/add", addHandler)        // Route to add a new To-Do item.
	http.HandleFunc("/complete", completeHandler) // Route to mark a To-Do item as completed.
	http.HandleFunc("/delete", deleteHandler)  // Route to delete a To-Do item.

	// Start the web server on port 8080.
	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// indexHandler handles the root route ("/") and displays the To-Do list.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the HTML template file.
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Execute the template and pass the list of todos to it.
	tmpl.Execute(w, todos)
}

// addHandler handles the "/add" route and adds a new To-Do item.
func addHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Get the title of the new task from the form data.
		title := r.FormValue("title")
		// Add the new task to the todos slice.
		todos = append(todos, Todo{ID: nextID, Title: title, Completed: false})
		// Increment the nextID to ensure each task has a unique ID.
		nextID++
	}
	// Redirect the user back to the root route ("/") after adding the task.
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// completeHandler handles the "/complete" route and marks a To-Do item as completed.
func completeHandler(w http.ResponseWriter, r *http.Request) {
	// Get the ID of the task to be marked as completed from the form data.
	id, _ := strconv.Atoi(r.FormValue("id"))
	// Find the task with the matching ID and mark it as completed.
	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Completed = true
			break
		}
	}
	// Redirect the user back to the root route ("/") after marking the task as completed.
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// deleteHandler handles the "/delete" route and deletes a To-Do item.
func deleteHandler(w http.ResponseWriter, r *http.Request) {
	// Get the ID of the task to be deleted from the form data.
	id, _ := strconv.Atoi(r.FormValue("id"))
	// Find the task with the matching ID and remove it from the todos slice.
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			break
		}
	}
	// Redirect the user back to the root route ("/") after deleting the task.
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
