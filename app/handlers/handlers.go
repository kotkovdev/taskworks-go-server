package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	tasks "taskworks/app/models/tasks"
	users "taskworks/app/models/users"
)

var fs = http.FileServer(nil)

func IndexHandler(w http.ResponseWriter, r *http.Request) {

}

func TasksHandler(w http.ResponseWriter, r *http.Request) {
	taskUser := users.User{Name: "Sergey Kotkov"}
	tasksList := &tasks.Task{Title: "Example task", User: taskUser}
	b, err := json.Marshal(tasksList)
	if err != nil {
		fmt.Println("error")
	}
	fmt.Fprintf(w, string(b))
}

func StaticHandler(routes []string) {
	fs = http.FileServer(http.Dir("./static/"))
	for _, route := range routes {
		http.Handle(route, http.StripPrefix(route, fs))
	}
}
