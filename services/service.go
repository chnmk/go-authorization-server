package service

import "fmt"

type User struct {
	Username     string `json:"username"`
	PasswordTemp string `json:"password"`
}

func service() {
	fmt.Println("Hello, World!")
}
