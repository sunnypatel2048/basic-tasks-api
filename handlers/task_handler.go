package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sunnypatel2048/basic-tasks-api/models"
)

var Tasks []models.Task

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Tasks)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range Tasks {
		if item.ID == params["id"] {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	errorResponse := map[string]string{"error": "Task not found"}
	json.NewEncoder(w).Encode(errorResponse)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	_ = json.NewDecoder(r.Body).Decode(&task)
	task.ID = fmt.Sprintf("%d", len(Tasks)+1)
	Tasks = append(Tasks, task)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var updatedTask models.Task
	_ = json.NewDecoder(r.Body).Decode(&updatedTask)

	for i, item := range Tasks {
		if item.ID == params["id"] {
			Tasks[i] = updatedTask
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedTask)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	errorResponse := map[string]string{"error": "Task not found"}
	json.NewEncoder(w).Encode(errorResponse)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, item := range Tasks {
		if item.ID == params["id"] {
			Tasks = append(Tasks[:i], Tasks[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	errorResponse := map[string]string{"error": "Task not found"}
	json.NewEncoder(w).Encode(errorResponse)
}
