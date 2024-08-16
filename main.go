package main

import (
	"fmt"
	"net/http"

	"github.com/rs/cors"

	transport "github.com/chnmk/sample-authorization-backend/transport/rest"
)

func main() {
	fmt.Println("App starting...")

	fmt.Println("Server starting...")
	mux := http.NewServeMux()
	mux.HandleFunc("/signup", transport.SignupHandler)
	mux.HandleFunc("/signin", transport.SigninHandler)

	// handlerDefault := cors.Default().Handler(mux)
	handler := cors.New(
		cors.Options{
			AllowedMethods:     []string{"POST", "OPTIONS"},
			AllowedOrigins:     []string{"*"},
			AllowedHeaders:     []string{"Bearer", "Bearer ", "authorization", "content-type", "authorization,content-type"},
			AllowCredentials:   true,
			OptionsPassthrough: true,
			Debug:              true,
		},
	).Handler(mux)

	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Shutting down...")
}
