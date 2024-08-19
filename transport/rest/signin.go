package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/chnmk/sample-authorization-backend/config"
	"github.com/golang-jwt/jwt"
)

// SigninHandler swagger:route POST /signin signinHandler
//
// Check if the submitted username and password exist and return a user permissions JWT.
func SigninHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("\nIncoming request to: /signin\n")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "authorization, content-type")

	if r.Method == http.MethodPost {
		fmt.Println("Got post request")
		header, user := readReq(w, r)

		// Read the authorization token from the request header
		if header != "" {
			token := strings.Split(header, " ")
			if len(token) == 2 {
				// Find the user in the database
				group, err := config.Database.Find(user.Username, token[1])
				if err != nil {
					fmt.Println(err)
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}

				// Return user's permission group and a confirmation token
				secret := []byte("authorization_changeme")

				claims := jwt.MapClaims{"group": group}
				token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

				signedToken, err := token.SignedString(secret)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				fmt.Println("User successfully found in the database")
				w.Write([]byte(signedToken))
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
