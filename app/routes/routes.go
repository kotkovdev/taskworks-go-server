package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"taskworks/app/handlers"
)

func Route() {
	r := mux.NewRouter().StrictSlash(true)

	// Handle API routes
	api := r.PathPrefix("/api/").Subrouter()
	api.HandleFunc("/tasks", handlers.TasksHandler)

	http.Handle("/", r)

}
