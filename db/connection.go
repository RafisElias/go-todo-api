package db

import (
	"database/sql"
	"fmt"
	"todo-api/configs"

	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()
	sc := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host,
		conf.Port,
		conf.User,
		conf.Password,
		conf.Database,
	)
	conn, err := sql.Open("postgres", sc)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	err = conn.Ping()
	return conn, err

}
