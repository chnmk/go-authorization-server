package handler

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request to: signup")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "authorization, content-type")

	//database := db.UseDB("default")

	if r.Method == http.MethodPost {
		fmt.Println("Got post request")

		// var user service.User
		header := r.Header.Get("Authorization")

		// Read authorization token from request header
		if header != "" {
			token := strings.Split(header, " ")
			if len(token) == 2 {
				// ...
				fmt.Println(token[0])
				fmt.Println(token[1])
				// database.add(username, token[1])

				w.Write([]byte("success"))
			} else {
				http.Error(w, "Invalid header", http.StatusBadRequest)
				return
			}
		} else {
			http.Error(w, "Invalid header", http.StatusBadRequest)
			return
		}

		// Read username and other data from rqeuest body
		body, err := io.ReadAll(r.Body)
		r.Body.Close()
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		fmt.Println(string(body))
		return
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
