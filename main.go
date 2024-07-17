package main

import (
	"fmt"
	"net/http"
)

func mainHandle(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Incoming request!")
	res.Write([]byte("Hello, World!"))
}

func main() {
	fmt.Println("App starting...")

	fmt.Println("Server starting...")
	http.HandleFunc(`/`, mainHandle)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("Shutting down...")
}
