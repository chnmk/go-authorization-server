package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/chnmk/sample-authorization-backend/config"
)

// SignupHandler swagger:route POST /signup signupHandler
//
// Save user credentials in a database if a user with this name doesn't already exist.
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("\nIncoming request to: /signup\n")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "authorization, content-type")

	if r.Method == http.MethodPost {
		fmt.Println("Got post request")
		header, user := readReq(w, r)

		// Read the authorization token from the request header
		if header != "" {
			token := strings.Split(header, " ")
			if len(token) == 2 {
				// Add a new user to the database
				err := config.Database.Add(user.Username, token[1], user.Group)
				if err != nil {
					fmt.Println(err)
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}

				fmt.Println("User successfully added to the database")
				w.Write([]byte("success"))
				return

			} else {
				http.Error(w, "Invalid header", http.StatusBadRequest)
				return
			}

		} else {
			http.Error(w, "Invalid header", http.StatusBadRequest)
			return
		}
	}

	handlePreflight(w, r)
}
