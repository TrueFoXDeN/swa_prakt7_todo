package controller

import (
	"html/template"
	"net/http"
	"strconv"
	"swa__prakt7_todo-03/app/model"
)

var tmpl *template.Template

// Is executed automatically on package load
func init() {
	tmpl = template.Must(template.ParseGlob("app/view/*.html"))
}

// Index controller
func Index(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	username := session.Values["username"].(string)

	todos, _ := model.GetAllTodosForUser(username)
	data := struct {
		Username string
		Todos    *[]model.Todo
	}{
		username,
		&todos,
	}
	tmpl.ExecuteTemplate(w, "todo.html", data)
}

// AddTodo controller
func AddTodo(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	username := session.Values["username"].(string)

	if r.Method == "POST" {
		action := r.FormValue("action")

		if len(action) != 0 {
			t := model.Todo{
				Action:   action,
				Done:     false,
				Username: username,
			}
			t.Add()
		}
	}

	todos, _ := model.GetAllTodosForUser(username)
	data := struct {
		Username string
		Todos    *[]model.Todo
	}{
		username,
		&todos,
	}
	tmpl.ExecuteTemplate(w, "todo.html", data)
}

// ToggleDone controller
func ToggleDone(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	username := session.Values["username"].(string)

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		data := struct {
			ErrorMsg string
		}{
			ErrorMsg: "Something was wrong with your request parameters!",
		}
		tmpl.ExecuteTemplate(w, "todo.html", data)
	}

	todo, _ := model.GetTodo(id, username)
	todo.ToggleStatus()

	todos, _ := model.GetAllTodosForUser(username)
	data := struct {
		Username string
		Todos    *[]model.Todo
	}{
		username,
		&todos,
	}
	tmpl.ExecuteTemplate(w, "todo.html", data)
}

// DeleteTodo controller
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	username := session.Values["username"].(string)

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		data := struct {
			ErrorMsg string
		}{
			ErrorMsg: "Something was wrong with your request parameters!",
		}
		tmpl.ExecuteTemplate(w, "todo.html", data)
	}

	todo, _ := model.GetTodo(id, username)
	todo.Delete()

	todos, _ := model.GetAllTodosForUser(username)
	data := struct {
		Username string
		Todos    *[]model.Todo
	}{
		username,
		&todos,
	}
	tmpl.ExecuteTemplate(w, "todo.html", data)
}
