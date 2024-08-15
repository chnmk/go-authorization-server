package database

import "github.com/chnmk/sample-authorization-backend/database/defaultDB"

type DB interface {
	Add(string, string) error
	Find(string, string) error
}

func UseDB(db string) DB {
	if db == "default" {
		var new defaultDB.DB
		return &new
	} else if db == "sqlite" {
		return nil
	} else {
		return nil
	}
}
