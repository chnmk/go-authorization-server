package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Struct for valid user data from transport/rest/read_test.go
type userValid struct {
	Username string `json:"username"`
	Group    string `json:"group"`
	Password string `json:"password"`
}

// Struct for invalid user data from transport/rest/read_test.go
type userInvalid struct {
	Invalid_data string `json:"invalid_data"`
	Group        string `json:"group"`
	Password     string `json:"password"`
}

func TestReadValidRequest(t *testing.T) {
	// Request data
	u := "user1"
	g := "admin"
	pass := "samplejwt"

	user := userValid{Username: u, Group: g}
	expected := fmt.Sprintf("Header: Bearer %s, body: {%s %s }", pass, u, g)

	// Encoded data
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(user)

	require.NoError(t, err)

	// Creating request
	req := httptest.NewRequest("POST", "/signin", &buf)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", pass))

	// Statring server
	rec := httptest.NewRecorder()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header, body := readReq(w, r)
		w.Write([]byte(fmt.Sprintf("Header: %s, body: %s", header, body)))
	})
	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, expected, rec.Body.String())
}

func TestReadValidRequestNoHeader(t *testing.T) {
	// Request data
	u := "user1"
	g := "admin"

	user := userValid{Username: u, Group: g}
	expected := fmt.Sprintf("Header: , body: {%s %s }", u, g)

	// Encoded data
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(user)

	require.NoError(t, err)

	// Creating request
	req := httptest.NewRequest("POST", "/signin", &buf)

	// Statring server
	rec := httptest.NewRecorder()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header, body := readReq(w, r)
		w.Write([]byte(fmt.Sprintf("Header: %s, body: %s", header, body)))
	})
	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, expected, rec.Body.String())
}

func TestReadInvalidRequest(t *testing.T) {
	// Request data
	u := "user1"
	g := "admin"
	pass := "samplejwt"

	user := userInvalid{Invalid_data: u, Group: g}
	expected := fmt.Sprintf("Header: Bearer %s, body: { %s }", pass, g)

	// Encoded data
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(user)

	require.NoError(t, err)

	// Creating request
	req := httptest.NewRequest("POST", "/signin", &buf)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", pass))

	// Statring server
	rec := httptest.NewRecorder()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header, body := readReq(w, r)
		w.Write([]byte(fmt.Sprintf("Header: %s, body: %s", header, body)))
	})
	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, expected, rec.Body.String())
}
