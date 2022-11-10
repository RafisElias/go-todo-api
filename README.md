# TODO API

this is a simple todo api

## Packages being used
- [go-chi](https://github.com/go-chi/chi)
  - Chi is a lightweight, idiomatic and composable router for building Go HTTP services.
- [viper](github.com/spf13/viper)
  - Viper is a complete configuration solution for Go applications including 12-Factor apps.
- [pq](github.com/lib/pq)
  - A pure Go postgres driver for Go's `database/sql` package.
- [uuid](https://github.com/google/uuid)
  - The uuid package generates and inspects UUIDs based on RFC 4122 and DCE 1.1: Authentication and Security Services.

## Before start

### Config file
first create a `config.toml` with the follow itens:

```toml
[api]
port = "9000"

[database]
host = "your postgres database host"
user = "your postgres user"
password = "your postgres user password"
name = "your database name"
```

### SQL commands
run this commands in the postgres terminal

```sql
CREATE DATABASE api_todo;

CREATE USER user_todo;

ALTER USER user_todo WITH ENCRYPTED PASSWORD '1122';

GRANT ALL PRIVILEGES ON DATABASE api_todo TO user_todo;

CREATE TABLE todos (id SERIAL PRIMARY KEY, title VARCHAR, description TEXT, done BOOL);

GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA PUBLIC TO user_todo;

GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA PUBLIC TO user_todo;
```

## Commands

All commands are run from the root of the project, from a terminal:

| Command                 | Action                                                                                                             |
| :---------------------- | :----------------------------------------------------------------------------------------------                    |
| `go run main.go`        | Starts server at `localhost:9000`                                                                                  |
| `go mod tyde`           | Download all the dependencies that are required in your source files and update go.mod file with that dependency.  |
| `go test`               | Runs the test files                                                                                                |