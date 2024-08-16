package config

import "github.com/chnmk/sample-authorization-backend/database"

var Database = database.UseDB("default")
