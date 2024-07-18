package main

import (
	"fmt"
	"net/http"

	"github.com/rs/cors"
)

func mainHandle(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Incoming request to: main")
	res.Write([]byte("Hello from main!"))
}

func signUpPost(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Incoming request to: signup")
	if req.Method == http.MethodPost {
		fmt.Println("Got post request")
		// fmt.Println(req.Body)
	}
	res.Write([]byte("Hello from signup!"))
}

func main() {
	fmt.Println("App starting...")

	fmt.Println("Server starting...")
	mux := http.NewServeMux()
	mux.HandleFunc("/", mainHandle)
	mux.HandleFunc("/signup", signUpPost)

	handler := cors.Default().Handler(mux)
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Shutting down...")
}
