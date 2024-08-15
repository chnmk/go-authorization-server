package defaultDB

import "errors"

type DB struct {
	users map[string]string
}

func (db *DB) Add(name string, token string) error {
	for key := range db.users {
		if key == name {
			return errors.New("user with this name already exists")
		}
	}

	db.users[name] = token
	return nil
}

func (db DB) Find(name string, token string) error {
	for key, value := range db.users {
		if key != name || value != token {
			return errors.New("invalid username or password")
		}
	}

	return nil
}
