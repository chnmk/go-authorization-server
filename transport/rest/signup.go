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

	// Check request method
	if r.Method != http.MethodPost {
		handlePreflight(w, r)
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

	// Add new user to the database
	err := config.Database.Add(user.Username, token[1], user.Group)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("User successfully added to the database")
	w.Write([]byte("success"))

}
