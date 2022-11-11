package todo

import "github.com/go-chi/chi/v5"

func Routers(r chi.Router) {
	r.Post("/", createTodoHandler)
	r.Get("/", findAllTodosHandler)
	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", findTodoByIdHandler)
		r.Put("/", updateTodoHandler)
		r.Delete("/", deleteTodoHandler)
	})
}
