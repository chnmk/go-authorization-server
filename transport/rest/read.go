package handler

import (
	"bytes"
	"encoding/json"
	"net/http"

	service "github.com/chnmk/sample-authorization-backend/services"
)

func readReq(w http.ResponseWriter, r *http.Request) (string, service.User) {
	header := r.Header.Get("Authorization")

	// Read request body
	var user service.User
	var buf bytes.Buffer

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return "", user
	}

	if err = json.Unmarshal(buf.Bytes(), &user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return "", user
	}

	return header, user
}
