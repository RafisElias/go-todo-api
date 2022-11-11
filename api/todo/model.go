package todo

import "errors"

var (
	ErrTodoIsEmpty           = errors.New("error: Todo can't be empty")
	ErrTitleIsRequired       = errors.New("error: title is required")
	ErrDescriptionIsRequired = errors.New("error: description is required")
)

type Todo struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

func (t *Todo) IsValid() error {
	if *t == (Todo{}) {
		return ErrTodoIsEmpty
	}

	if t.Title == "" {
		return ErrTitleIsRequired
	}

	if t.Description == "" {
		return ErrDescriptionIsRequired
	}

	return nil
}
