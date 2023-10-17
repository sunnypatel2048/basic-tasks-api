package main

import (
	"log"
	"net/http"

	"github.com/sunnypatel2048/basic-tasks-api/api"
)

func main() {
	router := api.SetupRoutes() // Initialize the router with your routes

	log.Fatal(http.ListenAndServe(":8080", router))
}
