package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/chnmk/sample-authorization-backend/config"
	"github.com/chnmk/sample-authorization-backend/database"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request to: signup")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "authorization, content-type")

	database := database.UseDB(config.DBConfig)

	if r.Method == http.MethodPost {
		fmt.Println("Got post request")
		header, user := readReq(w, r)

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

	handlePreflight(w, r)
}
