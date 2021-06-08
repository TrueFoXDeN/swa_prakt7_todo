package main

import (
	"net/http"
	"swa__prakt7_todo-03/app/controller"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// URL endpoints for HTML frontend
	r.HandleFunc("/", controller.Auth(controller.Index)).Methods("GET")
	r.HandleFunc("/add-todo", controller.Auth(controller.AddTodo)).Methods("POST")
	r.HandleFunc("/toggle-done", controller.Auth(controller.ToggleDone)).Methods("GET")
	r.HandleFunc("/delete-todo", controller.Auth(controller.DeleteTodo)).Methods("GET")

	// User regitration and login
	r.HandleFunc("/register", controller.Register).Methods("GET")
	r.HandleFunc("/add-user", controller.AddUser).Methods("POST")
	r.HandleFunc("/login", controller.Login).Methods("GET")
	r.HandleFunc("/authenticate-user", controller.AuthenticateUser).Methods("POST")
	r.HandleFunc("/logout", controller.Logout).Methods("GET")

	// Static resources
	r.PathPrefix("/css/").Handler(http.StripPrefix("/css", http.FileServer(http.Dir("./static/css"))))
	r.PathPrefix("/fonts/").Handler(http.StripPrefix("/fonts", http.FileServer(http.Dir("./static/fonts"))))
	r.PathPrefix("/js/").Handler(http.StripPrefix("/js", http.FileServer(http.Dir("./static/js"))))

	server := http.Server{
		Addr:    ":8383",
		Handler: r,
	}

	server.ListenAndServe()
}
