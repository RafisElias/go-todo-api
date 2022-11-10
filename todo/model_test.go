package todo

import (
	"testing"
)

func TestTodoModel(t *testing.T) {
	t.Run("should return a 'ErrTodoIsEmpty' when create a empty todo", func(t *testing.T) {
		var todo Todo
		result := todo.IsValid()
		if result != ErrTodoIsEmpty {
			t.Errorf("era esperado %v", ErrTodoIsEmpty)
		}
	})

	t.Run("should return a 'ErrTitleIsRequired' when the todo title is empty", func(t *testing.T) {
		todo := Todo{
			Description: "adasd",
		}

		result := todo.IsValid()
		if result != ErrTitleIsRequired {
			t.Errorf("era esperado %v", ErrTitleIsRequired)
		}
	})
}
