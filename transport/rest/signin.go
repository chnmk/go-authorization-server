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
// Check if submitted username and password exist and return user permissions JWT.
func SigninHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request to: signin")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "authorization, content-type")

	if r.Method == http.MethodPost {
		fmt.Println("Got post request")
		header, user := readReq(w, r)

		// Read authorization token from request header
		if header != "" {
			token := strings.Split(header, " ")
			if len(token) == 2 {
				// Find user in database
				group, err := config.Database.Find(user.Username, token[1])
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}

				// Return user group and confirmation token
				secret := []byte("authorization_changeme")

				claims := jwt.MapClaims{"group": group}
				token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

				signedToken, err := token.SignedString(secret)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

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
