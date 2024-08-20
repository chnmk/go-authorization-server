package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSigninValidRequest(t *testing.T) {
	// Sign up request data
	u := "TestSigninValidRequest"
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

	// Statring sign up server
	rec := httptest.NewRecorder()

	handler := http.HandlerFunc(SignupHandler)
	handler.ServeHTTP(rec, req)

	// Sign in data
	user2 := userValid{Username: u}

	var buf2 bytes.Buffer
	err = json.NewEncoder(&buf2).Encode(user2)

	require.NoError(t, err)

	// Creating sign in request
	req2 := httptest.NewRequest("POST", "/signin", &buf2)
	req2.Header.Set("Authorization", fmt.Sprintf("Bearer %s", pass))

	// Statring sign in server
	rec2 := httptest.NewRecorder()

	handler2 := http.HandlerFunc(SigninHandler)
	handler2.ServeHTTP(rec2, req2)

	require.Equal(t, http.StatusOK, rec2.Code)
	require.NotEmpty(t, rec2.Body.String())

	// Validate response JWT
	secret := []byte("authorization_changeme")
	_, err = jwt.Parse(rec2.Body.String(), func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	assert.NoError(t, err)
}

func TestSigninInvalidMethod(t *testing.T) {
	// Sign up request data
	u := "TestSigninInvalidMethod"
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

	// Statring sign up server
	rec := httptest.NewRecorder()

	handler := http.HandlerFunc(SignupHandler)
	handler.ServeHTTP(rec, req)

	// Sign in data
	user2 := userValid{Username: u}

	var buf2 bytes.Buffer
	err = json.NewEncoder(&buf2).Encode(user2)

	require.NoError(t, err)

	// Creating sign in request
	req2 := httptest.NewRequest("DELETE", "/signin", &buf2)
	req2.Header.Set("Authorization", fmt.Sprintf("Bearer %s", pass))

	// Statring sign in server
	rec2 := httptest.NewRecorder()

	handler2 := http.HandlerFunc(SigninHandler)
	handler2.ServeHTTP(rec2, req2)

	require.Equal(t, http.StatusMethodNotAllowed, rec2.Code)
	assert.Equal(t, "Method not allowed\n", rec2.Body.String())
}

func TestSigninInvalidUsername(t *testing.T) {
	// Sign up request data
	u := "TestSigninInvalidUsername"
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

	// Statring sign up server
	rec := httptest.NewRecorder()

	handler := http.HandlerFunc(SignupHandler)
	handler.ServeHTTP(rec, req)

	// Sign in data
	user2 := userValid{Username: u}

	var buf2 bytes.Buffer
	err = json.NewEncoder(&buf2).Encode(user2)

	require.NoError(t, err)

	// Creating sign in request
	req2 := httptest.NewRequest("POST", "/signin", &buf2)
	req2.Header.Set("Authorization", fmt.Sprintf("Bearer %s", pass))

	// Statring sign in server
	rec2 := httptest.NewRecorder()

	handler2 := http.HandlerFunc(SigninHandler)
	handler2.ServeHTTP(rec2, req2)

	require.Equal(t, http.StatusBadRequest, rec2.Code)
	assert.Equal(t, "invalid username or password\n", rec2.Body.String())
}

func TestSigninInvalidPassword(t *testing.T) {
	// Sign up request data
	u := "TestSigninInvalidPassword"
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

	// Statring sign up server
	rec := httptest.NewRecorder()

	handler := http.HandlerFunc(SignupHandler)
	handler.ServeHTTP(rec, req)

	// Sign in data
	pass2 := "wrongjwt"
	user2 := userValid{Username: u}

	var buf2 bytes.Buffer
	err = json.NewEncoder(&buf2).Encode(user2)

	require.NoError(t, err)

	// Creating sign in request
	req2 := httptest.NewRequest("POST", "/signin", &buf2)
	req2.Header.Set("Authorization", fmt.Sprintf("Bearer %s", pass2))

	// Statring sign in server
	rec2 := httptest.NewRecorder()

	handler2 := http.HandlerFunc(SigninHandler)
	handler2.ServeHTTP(rec2, req2)

	require.Equal(t, http.StatusBadRequest, rec2.Code)
	assert.Equal(t, "invalid username or password\n", rec2.Body.String())
}

func TestSigninUserDoesntExist(t *testing.T) {
	// Sign up request data
	u := "TestSigninUserDoesntExist"
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

	// Statring sign up server
	rec := httptest.NewRecorder()

	handler := http.HandlerFunc(SignupHandler)
	handler.ServeHTTP(rec, req)

	// Sign in data
	u2 := "TestSigninUserDoesntExistFakeUser"
	user2 := userValid{Username: u2}

	var buf2 bytes.Buffer
	err = json.NewEncoder(&buf2).Encode(user2)

	require.NoError(t, err)

	// Creating sign in request
	req2 := httptest.NewRequest("POST", "/signin", &buf2)
	req2.Header.Set("Authorization", fmt.Sprintf("Bearer %s", pass))

	// Statring sign in server
	rec2 := httptest.NewRecorder()

	handler2 := http.HandlerFunc(SigninHandler)
	handler2.ServeHTTP(rec2, req2)

	require.Equal(t, http.StatusBadRequest, rec2.Code)
	assert.Equal(t, "invalid username or password\n", rec2.Body.String())
}

func TestSigninEmptyHeader(t *testing.T) {
	// Sign up request data
	u := "TestSigninEmptyHeader"
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

	// Statring sign up server
	rec := httptest.NewRecorder()

	handler := http.HandlerFunc(SignupHandler)
	handler.ServeHTTP(rec, req)

	// Sign in data
	user2 := userValid{Username: u}

	var buf2 bytes.Buffer
	err = json.NewEncoder(&buf2).Encode(user2)

	require.NoError(t, err)

	// Creating sign in request
	req2 := httptest.NewRequest("POST", "/signin", &buf2)

	// Statring sign in server
	rec2 := httptest.NewRecorder()

	handler2 := http.HandlerFunc(SigninHandler)
	handler2.ServeHTTP(rec2, req2)

	require.Equal(t, http.StatusBadRequest, rec2.Code)
	assert.Equal(t, "Invalid header\n", rec2.Body.String())
}

func TestSigninInvalidHeaderNoBearer(t *testing.T) {
	// Sign up request data
	u := "TestSigninInvalidHeaderNoBearer"
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

	// Statring sign up server
	rec := httptest.NewRecorder()

	handler := http.HandlerFunc(SignupHandler)
	handler.ServeHTTP(rec, req)

	// Sign in data
	user2 := userValid{Username: u}

	var buf2 bytes.Buffer
	err = json.NewEncoder(&buf2).Encode(user2)

	require.NoError(t, err)

	// Creating sign in request
	req2 := httptest.NewRequest("POST", "/signin", &buf2)
	req2.Header.Set("Authorization", fmt.Sprintf("Bearer %s", pass))

	// Statring sign in server
	rec2 := httptest.NewRecorder()

	handler2 := http.HandlerFunc(SigninHandler)
	handler2.ServeHTTP(rec2, req2)

	require.Equal(t, http.StatusBadRequest, rec2.Code)
	assert.Equal(t, "invalid username or password\n", rec2.Body.String())
}

func TestSigninInvalidHeaderExtraData(t *testing.T) {
	// Sign up request data
	u := "TestSigninInvalidHeaderExtraData"
	g := "admin"
	pass := "samplejwt"

	user := userValid{Username: u, Group: g}

	// Encoded sign up data
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(user)

	require.NoError(t, err)

	// Creating sign up request
	req := httptest.NewRequest("POST", "/signup", &buf)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s Extra", pass))

	// Statring sign up server
	rec := httptest.NewRecorder()

	handler := http.HandlerFunc(SignupHandler)
	handler.ServeHTTP(rec, req)

	// Sign in data
	user2 := userValid{Username: u}

	var buf2 bytes.Buffer
	err = json.NewEncoder(&buf2).Encode(user2)

	require.NoError(t, err)

	// Creating sign in request
	req2 := httptest.NewRequest("POST", "/signin", &buf2)
	req2.Header.Set("Authorization", fmt.Sprintf("Bearer %s", pass))

	// Statring sign in server
	rec2 := httptest.NewRecorder()

	handler2 := http.HandlerFunc(SigninHandler)
	handler2.ServeHTTP(rec2, req2)

	require.Equal(t, http.StatusBadRequest, rec2.Code)
	assert.Equal(t, "invalid username or password\n", rec2.Body.String())
}
