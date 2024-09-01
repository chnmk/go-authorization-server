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

	// Check request method
	if r.Method != http.MethodPost {
		handlePreflight(w, r)
		return
	}
	fmt.Println("Got post request")

	// Try to read the request
	header, user := readReq(w, r)
	if header == "" {
		http.Error(w, "Invalid header", http.StatusBadRequest)
		return
	}

	// Confirm that the request header contains a JWT Bearer token
	token := strings.Split(header, " ")
	if len(token) != 2 {
		http.Error(w, "Invalid header", http.StatusBadRequest)
		return
	}

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
	newJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := newJWT.SignedString(secret)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("User successfully found in the database")
	w.Write([]byte(signedToken))
}
