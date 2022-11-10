package todo

import (
	"database/sql"
	"fmt"

	"github.com/RafisElias/todo-api/db"
)

func InsertTodo(todo Todo) (insertedTodo Todo, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	err = conn.QueryRow(`
		INSERT INTO todos (title, description, done)
		VALUES ($1, $2, $3) 
		RETURNING id, title, description, done`,
		todo.Title, todo.Description, todo.Done,
	).Scan(
		&insertedTodo.ID,
		&insertedTodo.Title,
		&insertedTodo.Description,
		&insertedTodo.Done,
	)

	return
}

func GetTodo(id int64) (Todo, error) {
	var todo Todo
	conn, err := db.OpenConnection()
	if err != nil {
		return todo, err
	}
	defer conn.Close()

	row := conn.QueryRow(`select * from todos where id = $1`, id)
	err = row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)
	return todo, err
}

func GetAllTodos() ([]Todo, error) {
	var todos []Todo
	conn, err := db.OpenConnection()
	if err != nil {
		return todos, err
	}
	defer conn.Close()

	rows, err := conn.Query("select * from todos")
	if err != nil {
		return todos, err
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)
		if err != nil {
			continue
		}
		todos = append(todos, todo)
	}

	return todos, err
}

func UpdateTodo(id int64, todo Todo) (Todo, error) {
	var updatedTodo Todo
	conn, err := db.OpenConnection()
	if err != nil {
		return updatedTodo, err
	}
	defer conn.Close()

	row := conn.QueryRow(`
		UPDATE todos SET title=$1, description=$2, done=$3
		WHERE id=$4
		RETURNING id, title, description, done
	`,
		todo.Title,
		todo.Description,
		todo.Done,
		id,
	)

	err = row.Scan(&updatedTodo.ID, &updatedTodo.Title, &updatedTodo.Description, &updatedTodo.Done)

	return updatedTodo, err
}

func DeleteTodo(id int64) error {
	conn, err := db.OpenConnection()
	if err != nil {
		return err
	}
	defer conn.Close()

	row := conn.QueryRow(`DELETE FROM todos WHERE id=$1 RETURNING id`, id)

	err = row.Scan()
	if err == sql.ErrNoRows {
		err = fmt.Errorf("todo with id '%d' not found", id)
	}

	return err
}
