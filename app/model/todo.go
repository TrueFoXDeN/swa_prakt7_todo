package model

import (
	"database/sql"
	"errors"
	"fmt"

	// blank import of postgresql database driver
	_ "github.com/lib/pq"
)

// Todo data structure
type Todo struct {
	ID       int
	Action   string
	Done     bool
	Username string
}

// DB handle
var db *sql.DB

// DB constanta
const (
	dbHost = "localhost"
	dbPort = 5432
	dbUser = "postgres"
	dbPwd  = "postgres"
	dbName = "todo"
)

func init() {
	var err error
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPwd, dbName)
	db, err = sql.Open("postgres", dbinfo)
	//defer db.Close()

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

// Add todo to DB
func (t Todo) Add() (int, error) {
	var lastInsertedID int
	err := db.QueryRow("INSERT INTO todos(action,done,username) VALUES($1,$2,$3) returning id;", t.Action, t.Done, t.Username).Scan(&lastInsertedID)

	if err != nil {
		fmt.Printf("[Add] ERROR: %s\n", err)
	}

	return lastInsertedID, err
}

// GetTodo retrieves todo with the provided id from DB
func GetTodo(id int, username string) (Todo, error) {
	var todo Todo
	q := fmt.Sprintf("SELECT id, action, done, username FROM todos WHERE id=%d and username='%s'", id, username)
	err := db.QueryRow(q).Scan(&todo.ID, &todo.Action, &todo.Done, &todo.Username)
	if err != nil {
		return Todo{}, errors.New("no todo found")
	}

	if todo.Username != username {
		return Todo{}, errors.New("not authorized")
	}

	return todo, nil
}

// GetAllTodosForUser retrieves Todos of provided username from DB
func GetAllTodosForUser(username string) ([]Todo, error) {
	rows, err := db.Query("SELECT * FROM todos WHERE username='" + username + "'")
	if err != nil {
		return nil, errors.New("no todo found")
	}
	defer rows.Close()

	var allTodos []Todo
	var tempTodo Todo
	for rows.Next() {
		err := rows.Scan(&tempTodo.ID, &tempTodo.Action, &tempTodo.Done, &tempTodo.Username)
		if err != nil {
			return nil, errors.New("error while retrieving all todos")
		}

		allTodos = append(allTodos, Todo{ID: tempTodo.ID, Action: tempTodo.Action, Done: tempTodo.Done, Username: tempTodo.Username})
	}
	if err := rows.Err(); err != nil {
		return nil, errors.New("error while retrieving all todos")
	}

	return allTodos, nil
}

// ToggleStatus changes the completion status of the Todo with the provided id contained in DB
func (t Todo) ToggleStatus() error {
	if t.Done {
		t.Done = false
	} else {
		t.Done = true
	}

	err := db.QueryRow("UPDATE todos SET done=$1 WHERE id=$2", t.Done, t.ID).Scan()

	if err != nil {
		fmt.Printf("[ToggleStatus] ERROR: %s\n", err)
	}

	return err
}

// Delete Todo with the provided id from DB
func (t Todo) Delete() error {
	sqlStatement := `
		DELETE FROM todos
		WHERE id = $1;`

	_, err := db.Exec(sqlStatement, t.ID)

	if err != nil {
		fmt.Printf("[Delete] ERROR: %s\n", err)
	}

	return err
}
