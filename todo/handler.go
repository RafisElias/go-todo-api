package todo

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type ErrorResponse struct {
	Message    string `json:"message"`
	statusCode int
}

func createTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	w.Header().Add("Content-type", "application/json")
	if err != nil {
		errorHandler(w, ErrorResponse{Message: err.Error(), statusCode: 400})
		return
	}

	err = todo.IsValid()

	if err != nil {
		errorHandler(w, ErrorResponse{Message: err.Error(), statusCode: 400})
		return
	}

	insertedTodo, err := InsertTodo(todo)
	if err != nil {
		errorHandler(w, ErrorResponse{Message: err.Error(), statusCode: 400})
		return
	}
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(insertedTodo)
}

func updateTodoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	w.Header().Add("Content-type", "application/json")
	if err != nil {
		errorHandler(w, ErrorResponse{Message: err.Error(), statusCode: 422})
		return
	}

	var todo Todo

	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		errorHandler(w, ErrorResponse{Message: err.Error(), statusCode: 400})
		return
	}

	err = todo.IsValid()

	if err != nil {
		errorHandler(w, ErrorResponse{Message: err.Error(), statusCode: 400})
		return
	}

	updatedTodo, err := UpdateTodo(int64(id), todo)

	if err != nil {
		errorHandler(w, ErrorResponse{Message: err.Error(), statusCode: 400})
		return
	}

	json.NewEncoder(w).Encode(updatedTodo)
}

func deleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	w.Header().Add("Content-type", "application/json")
	if err != nil {
		errorHandler(w, ErrorResponse{Message: err.Error(), statusCode: 422})
		return
	}

	err = DeleteTodo(int64(id))

	if err != nil {
		errorHandler(w, ErrorResponse{Message: err.Error(), statusCode: 404})
		return
	}

	w.WriteHeader(204)
}

func findAllTodosHandler(w http.ResponseWriter, r *http.Request) {
	todos, err := GetAllTodos()
	if err != nil {
		log.Println(err)
	}
	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func findTodoByIdHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	w.Header().Add("Content-type", "application/json")
	if err != nil {
		errorHandler(w, ErrorResponse{
			Message:    err.Error(),
			statusCode: 422,
		})
		return
	}

	todo, err := GetTodo(int64(id))

	if err != nil {
		errorHandler(w, ErrorResponse{
			Message:    fmt.Sprintf("Todo with id '%d' not found", id),
			statusCode: 404,
		})
		return
	}

	json.NewEncoder(w).Encode(todo)
}

func errorHandler(w http.ResponseWriter, err ErrorResponse) {
	w.WriteHeader(err.statusCode)
	json.NewEncoder(w).Encode(err)
}
