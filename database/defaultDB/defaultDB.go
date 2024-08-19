package defaultDB

import (
	"errors"
)

// Implements DB interface from database package
type DB struct {
	Users map[string][2]string
}

// Adds new user to the database. Returns error if user with this name already exists.
func (db *DB) Add(name string, token string, group string) error {
	for key := range db.Users {
		if key == name {
			return errors.New("user with this name already exists")
		}
	}

	db.Users[name] = [2]string{token, group}
	return nil
}

// Returns user permission group. Returns error if user with this name doesn't exist.
func (db DB) Find(name string, token string) (string, error) {
	var userExists bool
	var group string

	for key, value := range db.Users {
		if key == name && value[0] == token {
			userExists = true
			group = value[1]
		}
	}

	if !userExists {
		return "", errors.New("invalid username or password")
	}

	return group, nil
}
