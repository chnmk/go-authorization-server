package defaultDB

import "errors"

type DB struct {
	users map[string][2]string
}

func (db *DB) Add(name string, token string, group string) error {
	for key := range db.users {
		if key == name {
			return errors.New("user with this name already exists")
		}
	}

	db.users[name] = [2]string{token, group}
	return nil
}

func (db DB) Find(name string, token string) (string, error) {
	var userExists bool
	var group string

	for key, value := range db.users {
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
