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

func TestSignupValidRequest(t *testing.T) {
	// Sign up request data
	u := "user1"
	g := "admin"
	pass := "samplejwt"

	user := userValid{Username: u, Group: g}

	// Encoded sign up data
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(user)

	require.NoError(t, err)

	// Creating sign up request
	req := httptest.NewRequest("POST", "/signup", &buf)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", pass))

	// Statring server
	rec := httptest.NewRecorder()

	handler := http.HandlerFunc(SignupHandler)
	handler.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "success", rec.Body.String())
}

func TestSignupInvalidMethod(t *testing.T) {
	// Sign up request data
	u := "user1"
	g := "admin"
	pass := "samplejwt"

	user := userValid{Username: u, Group: g}

	// Encoded sign up data
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(user)

	require.NoError(t, err)

	// Creating sign up request
	req := httptest.NewRequest("GET", "/signup", &buf)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", pass))

	// Statring server
	rec := httptest.NewRecorder()

	handler := http.HandlerFunc(SignupHandler)
	handler.ServeHTTP(rec, req)

	require.Equal(t, http.StatusMethodNotAllowed, rec.Code)
}

func TestSignupInvalidBody(t *testing.T) {
	// Sign up request data
	u := "user1"
	g := "admin"
	pass := "samplejwt"

	user := userInvalid{Invalid_data: u, Group: g}

	// Encoded sign up data
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(user)

	require.NoError(t, err)

	// Creating sign up request
	req := httptest.NewRequest("POST", "/signup", &buf)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", pass))

	// Statring server
	rec := httptest.NewRecorder()

	handler := http.HandlerFunc(SignupHandler)
	handler.ServeHTTP(rec, req)

	require.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestSignupUserAlreadyExists(t *testing.T) {
	// Sign up request data
	u := "user1"
	g := "admin"
	pass := "samplejwt"

	// Create first user
	user := userValid{Username: u, Group: g}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(user)

	require.NoError(t, err)

	req := httptest.NewRequest("POST", "/signup", &buf)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", pass))

	rec := httptest.NewRecorder()

	handler := http.HandlerFunc(SignupHandler)
	handler.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "success", rec.Body.String())

	// Create second user
	user2 := userValid{Username: u, Group: g}

	var buf2 bytes.Buffer
	err = json.NewEncoder(&buf2).Encode(user2)

	require.NoError(t, err)

	req2 := httptest.NewRequest("POST", "/signup", &buf2)
	req2.Header.Set("Authorization", fmt.Sprintf("Bearer %s", pass))

	rec2 := httptest.NewRecorder()

	handler2 := http.HandlerFunc(SignupHandler)
	handler2.ServeHTTP(rec2, req2)

	// Assert result
	require.Equal(t, http.StatusBadRequest, rec2.Code)
	assert.Equal(t, "user with this name already exists\n", rec2.Body.String())
}

func TestSignupEmptyHeader(t *testing.T) {
	// Sign up request data
	u := "user1"
	g := "admin"
	// pass := "samplejwt"

	user := userValid{Username: u, Group: g}

	// Encoded sign up data
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(user)

	require.NoError(t, err)

	// Creating sign up request
	req := httptest.NewRequest("POST", "/signup", &buf)
	// req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", pass))

	// Statring server
	rec := httptest.NewRecorder()

	handler := http.HandlerFunc(SignupHandler)
	handler.ServeHTTP(rec, req)

	require.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "Invalid header\n", rec.Body.String())
}

func TestSignupInvalidHeaderNoBearer(t *testing.T) {
	// Sign up request data
	u := "user1"
	g := "admin"
	pass := "samplejwt"

	user := userValid{Username: u, Group: g}

	// Encoded sign up data
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(user)

	require.NoError(t, err)

	// Creating sign up request
	req := httptest.NewRequest("POST", "/signup", &buf)
	req.Header.Set("Authorization", pass)

	// Statring server
	rec := httptest.NewRecorder()

	handler := http.HandlerFunc(SignupHandler)
	handler.ServeHTTP(rec, req)

	require.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "Invalid header\n", rec.Body.String())
}

func TestSignupInvalidHeaderExtraData(t *testing.T) {
	// Sign up request data
	u := "user1"
	g := "admin"
	pass := "samplejwt"

	user := userValid{Username: u, Group: g}

	// Encoded sign up data
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(user)

	require.NoError(t, err)

	// Creating sign up request
	req := httptest.NewRequest("POST", "/signup", &buf)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s Extra_data", pass))

	// Statring server
	rec := httptest.NewRecorder()

	handler := http.HandlerFunc(SignupHandler)
	handler.ServeHTTP(rec, req)

	require.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "Invalid header\n", rec.Body.String())
}
