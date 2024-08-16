package handler

type User struct {
	Username string `json:"username"`
	Group    string `json:"group"`
}

const DBConfig = "default"
