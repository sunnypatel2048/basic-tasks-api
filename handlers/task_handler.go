package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sunnypatel2048/basic-tasks-api/models"
)

var tasks []models.Task

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range tasks {
		if item.ID == params["id"] {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	_ = json.NewDecoder(r.Body).Decode(&task)
	task.ID = fmt.Sprintf("%d", len(tasks)+1)
	tasks = append(tasks, task)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, item := range tasks {
		if item.ID == params["id"] {
			var task models.Task
			_ = json.NewDecoder(r.Body).Decode(&task)
			task.ID = params["id"]
			tasks[i] = task
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, item := range tasks {
		if item.ID == params["id"] {
			tasks = append(tasks[:i], tasks[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}
