package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func readReq(w http.ResponseWriter, r *http.Request) (string, User) {
	header := r.Header.Get("Authorization")

	// Read request body
	var user User
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
