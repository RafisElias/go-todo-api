package todo

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetTodos(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/todos", nil)

	fmt.Print(req.URL.Scheme)

	if err != nil {
		t.Errorf("Error try ger the todos: %v", err)
	}
}
