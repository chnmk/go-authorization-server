package handler

import (
	"fmt"
	"net/http"
	"strings"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request to: signup")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	if r.Method == http.MethodPost {
		fmt.Println("Got post request")

		// var user service.User
		header := r.Header.Get("Authorization")

		if header != "" {
			token := strings.Split(header, " ")
			if len(token) == 2 {
				// ...
				fmt.Println(token[0])
				fmt.Println(token[1])

				w.Write([]byte("success"))
				return
			}
		}
		http.Error(w, "Header error", http.StatusBadRequest)
		return
	}

	if r.Method != http.MethodOptions {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
