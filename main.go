package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"todo-api/api"
	"todo-api/api/todo"

	"todo-api/configs"
	m "todo-api/middleware"

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

	r.Get("/", IndexHandler)
	r.Handle(
		"/css/*",
		http.StripPrefix(
			"/css/",
			http.FileServer(http.Dir("css")),
		),
	)

	r.Route("/api", api.Routers)
	r.NotFound(notFoundRoute)

	fmt.Printf("⚡️[server]: Server is running at http://localhost%s \n", PORT)
	if err := http.ListenAndServe(PORT, r); err != nil {
		log.Panic(err)
	}
}

func notFoundRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	if strings.Contains(r.Header.Get("Accept"), "text/html") {
		if err := temp.ExecuteTemplate(w, "NotFound", nil); err != nil {
			log.Println(err)
		}
	} else if strings.Contains(r.Header.Get("Accept"), "json") {
		w.Header().Add("Content-type", "application/json")
		if err := json.NewEncoder(w).Encode(map[string]string{
			"message": fmt.Sprintf(`🔍 - Not Found - %s`, r.URL),
		}); err != nil {
			log.Println(err)
		}
	} else {
		w.Header().Add("Content-type", "text/plain")
		if _, err := w.Write(
			[]byte(fmt.Sprintf(`🔍 - Not Found - %s`, r.URL)),
		); err != nil {
			log.Println(err)
		}
	}
}
