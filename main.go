package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rs/cors"
)

type User struct {
	Username     string `json:"username"`
	PasswordTemp string `json:"password"`
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request to: main")
	w.Write([]byte("Hello from main!"))
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request to: signup")

	if r.Method == http.MethodPost {
		fmt.Println("Got post request")

		var user User
		var buf bytes.Buffer

		_, err := buf.ReadFrom(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err = json.Unmarshal(buf.Bytes(), &user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Println("Got username: ", user.Username)

		// TODO: change response writer
		w.Write([]byte("Hello from signup!"))
	}
}

func main() {
	fmt.Println("App starting...")

	fmt.Println("Server starting...")
	mux := http.NewServeMux()
	mux.HandleFunc("/", mainHandler)
	mux.HandleFunc("/signup", signupHandler)

	handler := cors.Default().Handler(mux)
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Shutting down...")
}
