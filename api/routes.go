package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sunnypatel2048/basic-tasks-api/handlers"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// Define API routes
	router.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", handlers.GetTask).Methods("GET")
	router.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE")

	// Handle 404 Not Found
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Found"))
	})

	return router
}
