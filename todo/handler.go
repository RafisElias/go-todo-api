package todo

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func createTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	w.Header().Add("Content-type", "application/json")
	if err != nil {
		errorHandler(w, 400, err)
		return
	}

	err = todo.IsValid()

	if err != nil {
		errorHandler(w, 400, err)
		return
	}

	insertedTodo, err := InsertTodo(todo)
	if err != nil {
		errorHandler(w, 400, err)
		return
	}
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(insertedTodo)
}

func updateTodoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	w.Header().Add("Content-type", "application/json")
	if err != nil {
		errorHandler(w, 422, err)
		return
	}

	var todo Todo

	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		errorHandler(w, 400, err)
		return
	}

	updatedTodo, err := UpdateTodo(int64(id), todo)

	log.Print(updatedTodo, err)
	if err != nil {
		errorHandler(w, 400, err)
		return
	}

	json.NewEncoder(w).Encode(updatedTodo)
}

func deleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	w.Header().Add("Content-type", "application/json")
	if err != nil {
		errorHandler(w, 422, err)
		return
	}

	err = DeleteTodo(int64(id))

	if err != nil {
		errorHandler(w, 404, err)
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
		errorHandler(w, 422, err)
		return
	}

	todo, err := GetTodo(int64(id))

	if err != nil {
		errorHandler(w, 404, fmt.Errorf("Todo with id '%d' not found", id))
		return
	}

	json.NewEncoder(w).Encode(todo)
}

func errorHandler(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{
		"message": fmt.Sprint(err),
	})
}
