# Basic Tasks API using Go

This is a simple REST API project implemented in Go. It provides basic functionality for managing tasks.

## Table of Contents
- [Getting Started](#getting-started)
- [Project Structure](#project-structure)
- [Endpoints](#endpoints)
- [Usage](#usage)
- [Dockerization](#dockerization)

## Getting Started

To get started with this project, you'll need Go installed on your system. Clone this repository and install the necessary dependencies using:

```bash
go mod tidy
```

## Project Structure

The project is structured as follows:
```lua
|-- basic-tasks-api/
|    |-- main.go
|    |-- go.mod
|    |-- go.sum
|    |-- api/
|    |   |-- routes.go
|    |-- models/
|    |   |-- task.go
|    |-- handlers/
|    |   |-- task_handler.go
```

- main.go: The entry point of the application.
- api/routes.go: Defines the API routes and handlers.
- models/task.go: Defines the Task data structure.
- handlers/task_handler.go: Contains the HTTP request handlers for tasks.

## Endpoints

- GET /tasks: Get a list of tasks.
- GET /tasks/{id}: Get a specific task by ID.
- POST /tasks: Create a new task.
- PUT /tasks/{id}: Update a task by ID.
- DELETE /tasks/{id}: Delete a task by ID.

## Usage

1. Start the application by running the following command in the project root directory:
    ```bash
    go run main.go
    ```

2. The API server will start, and you can use tools like curl, Postman, or your web browser to interact with the endpoints.

3. Example usage:

    - Create a new task:
        ```bash
        curl -X POST -d '{"title": "Task Title"}' http://localhost:8080/tasks
        ```

    - Get a list of tasks:
        ```bash
        curl http://localhost:8080/tasks
        ```

    - Get a specific task by ID:
        ```bash
        curl http://localhost:8080/tasks/1
        ```

    - Update a task:
        ```bash
        curl -X PUT -d '{"title": "Updated Task Title"}' http://localhost:8080/tasks/1
        ```

    - Delete a task:
        ```bash
        curl -X DELETE http://localhost:8080/tasks/1
        ```

## Dockerization

You can also run this project in a Docker container. Follow these steps:

1. Build a Docker image from the project root directory:
    ```bash
    docker build -t myapi .
    ```
    Replace myapi with your preferred image name.

2. Run the Docker container:
    ```bash
    docker run -p 8080:8080 myapi
    ```
    This maps port 8080 from the container to the host.

3. Access the application in your browser at http://localhost:8080.