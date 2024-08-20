package database

import "github.com/chnmk/sample-authorization-backend/database/defaultDB"

type DB interface {
	// Adds new user to the database. Returns error if user with this name already exists.
	//
	// Add(username, token, group) error
	Add(string, string, string) error

	// Returns user permission group. Returns error if user with this name doesn't exist.
	//
	// Find(username, token) (group, error)
	Find(string, string) (string, error)
}

// Sets current database based on a string value passed from the config package
func UseDB(db string) DB {
	if db == "sqlite" {
		panic("Not implemented yet!")
	} else { // db == "default" and every invalid input
		var new defaultDB.DB
		new.Users = make(map[string][2]string)
		return &new
	}
}
