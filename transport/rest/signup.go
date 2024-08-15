package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/chnmk/sample-authorization-backend/database"
)

type User struct {
	Username string `json:"username"`
	Group    string `json:"group"`
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request to: signup")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "authorization, content-type")

	database := database.UseDB("default")

	if r.Method == http.MethodPost {
		fmt.Println("Got post request")
		header := r.Header.Get("Authorization")

		// Read request body
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

		// Read authorization token from request header
		if header != "" {
			token := strings.Split(header, " ")
			if len(token) == 2 {
				// Add new user to database
				err := database.Add(user.Username, token[1], user.Group)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				w.Write([]byte("success"))

			} else {
				http.Error(w, "Invalid header", http.StatusBadRequest)
				return
			}

		} else {
			http.Error(w, "Invalid header", http.StatusBadRequest)
			return
		}
	}

	// Handle preflight requests
	if r.Method != http.MethodOptions {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	} else {
		w.Write([]byte("success"))
		return
	}
}
