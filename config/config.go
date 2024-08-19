package config

import "github.com/chnmk/sample-authorization-backend/database"

/*
	Any DB from the database package:

	database/defaultDB ("default")

	database/sqlite ("sqlite") - Work in progress

	database/...
*/
var Database = database.UseDB("default")
