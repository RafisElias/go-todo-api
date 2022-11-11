package api

import (
	"github.com/go-chi/chi/v5"
	"todo-api/api/todo"
)

func Routers(r chi.Router) {
	r.Route("/todos", todo.Routers)
}
