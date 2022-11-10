package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/RafisElias/todo-api/configs"
	m "github.com/RafisElias/todo-api/middleware"
	"github.com/RafisElias/todo-api/todo"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Index", nil)
}

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}
	var PORT = fmt.Sprintf(":%s", configs.GetServerPort())

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(m.RequestLogger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "Index", nil)
	})

	r.Handle(
		"/css/*",
		http.StripPrefix(
			"/css/",
			http.FileServer(http.Dir("css")),
		),
	)

	r.Route("/api/todos", todo.TodoRouters)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		if strings.Contains(r.Header.Get("Accept"), "text/html") {
			temp.ExecuteTemplate(w, "NotFound", nil)
		} else if strings.Contains(r.Header.Get("Accept"), "json") {
			w.Header().Add("Content-type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": fmt.Sprintf(`🔍 - Not Found - %s`, r.URL),
			})
		} else {
			w.Header().Add("Content-type", "text/plain")
			w.Write([]byte(fmt.Sprintf(`🔍 - Not Found - %s`, r.URL)))
		}
	})

	fmt.Printf("⚡️[server]: Server is running at http://localhost%s \n", PORT)
	http.ListenAndServe(PORT, r)
}
