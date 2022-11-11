package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
)

// RequestLogger HTTP middleware setting a value on the request context
func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		message := fmt.Sprintf("%s\t%s\t%s", r.Method, r.URL, r.Host)
		logEvents(message, "reqLog.log")
		next.ServeHTTP(w, r)
	})
}

func logEvents(message, filename string) {
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
		dateTime := time.Now()
		message = fmt.Sprintf(
			"%s\t%v\t%s\n",
			dateTime,
			uuid.New(),
			message,
		)
		_, err = file.Write([]byte(message))

		if err != nil {
			log.Print(err)
		}
	}

	defer file.Close()
}
