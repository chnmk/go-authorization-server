package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	service "github.com/chnmk/sample-authorization-backend/services"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request to: signup")

	if r.Method == http.MethodPost {
		fmt.Println("Got post request")

		var user service.User
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
