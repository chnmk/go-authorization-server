package main

import (
	"fmt"
	"net/http"
)

func mainHandle(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Incoming request to: main")
	res.Write([]byte("Hello from main!"))
}

func signUpPost(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Incoming request to: signup")
	if req.Method == http.MethodPost {
		fmt.Println("Got post request")
	}
	res.Write([]byte("Hello from signup!"))
}

func main() {
	fmt.Println("App starting...")

	fmt.Println("Server starting...")
	http.HandleFunc(`/`, mainHandle)
	http.HandleFunc(`/signup`, signUpPost)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("Shutting down...")
}
