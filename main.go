package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/RafisElias/todo-api/configs"
	"github.com/RafisElias/todo-api/todo"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
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
	r.Use(requestLogger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "Index", nil)
	})

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

// HTTP middleware setting a value on the request context
func requestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := os.MkdirAll("logs", 0750)
		if err != nil && !os.IsExist(err) {
			log.Print(err)
		}

		file, err := os.OpenFile(
			"logs/reqLog.log",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY,
			0666,
		)

		if err != nil {
			log.Print(err)
		} else {
			dateTime := time.Now().Format("20060102 15:04:05")
			id := uuid.New()
			message := fmt.Sprintf("%s\t%v\t%s\t%s\t%s\n", dateTime, id, r.Method, r.URL, r.Host)
			_, err = file.Write([]byte(message))

			if err != nil {
				log.Print(err)
			}
		}

		defer file.Close()

		next.ServeHTTP(w, r)
	})
}
