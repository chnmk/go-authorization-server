package database

import "github.com/chnmk/sample-authorization-backend/database/defaultDB"

type DB interface {
	// (username, token, group) error
	Add(string, string, string) error

	// (username, token) (group, error)
	Find(string, string) (string, error)
}

func UseDB(db string) DB {
	if db == "default" {
		var new defaultDB.DB
		new.Users = make(map[string][2]string)
		return &new
	} else if db == "sqlite" {
		return nil
	} else {
		return nil
	}
}
