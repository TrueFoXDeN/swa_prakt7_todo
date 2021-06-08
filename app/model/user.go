package model

import (
	"errors"
	"fmt"
)

// User data structure
type User struct {
	Username string
	Password string
}

// Add User
func (user User) Add() (err error) {
	// Check wether username already exists
	userInDB, err := GetUserByUsername(user.Username)
	if err == nil && userInDB.Username == user.Username {
		return errors.New("username exists already")
	}

	// Add user to DB
	_, err = db.Query("INSERT INTO users(username,password) VALUES($1,$2)", user.Username, user.Password)

	if err != nil {
		fmt.Printf("[Add] error: %s\n", err)
	}

	return err
}

// GetUserByUsername retrieve User by username
func GetUserByUsername(username string) (user User, err error) {
	if username == "" {
		return User{}, errors.New("no username provided")
	}

	stmt := "SELECT username, password FROM users WHERE username='" + username + "'"
	var dbUsername string
	var dbPassword string
	err = db.QueryRow(stmt).Scan(&dbUsername, &dbPassword)
	if err != nil || username != dbUsername {
		return User{}, err
	}

	user = User{dbUsername, dbPassword}

	return user, nil
}
