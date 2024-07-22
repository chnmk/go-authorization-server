package main

import (
	"fmt"
	"net/http"

	"github.com/rs/cors"

	handler "github.com/chnmk/sample-authorization-backend/transport/rest"
)

func main() {
	fmt.Println("App starting...")

	fmt.Println("Server starting...")
	mux := http.NewServeMux()
	mux.HandleFunc("/signup", handler.SignupHandler)

	handler := cors.Default().Handler(mux)
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Shutting down...")
}
