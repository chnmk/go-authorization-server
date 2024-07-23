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

	//handlerDefault := cors.Default().Handler(mux)
	handler := cors.New(
		cors.Options{
			AllowedMethods:     []string{"POST", "OPTIONS"},
			AllowedOrigins:     []string{"*"},
			AllowedHeaders:     []string{"Bearer", "Bearer ", "authorization", "content-type", "authorization,content-type"},
			AllowCredentials:   true,
			OptionsPassthrough: true,
			// Enable Debugging for testing, consider disabling in production
			Debug: true,
		},
	).Handler(mux)

	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Shutting down...")
}
