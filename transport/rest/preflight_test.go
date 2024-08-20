package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignupPreflightInvalidMethod(t *testing.T) {
	req := httptest.NewRequest("DELETE", "/signup", nil)
	rec := httptest.NewRecorder()

	handler := http.HandlerFunc(SignupHandler)
	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusMethodNotAllowed, rec.Code)
	assert.Equal(t, "Method not allowed\n", rec.Body.String())
}

func TestSigninPreflightInvalidMethod(t *testing.T) {
	req := httptest.NewRequest("DELETE", "/signin", nil)
	rec := httptest.NewRecorder()

	handler := http.HandlerFunc(SigninHandler)
	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusMethodNotAllowed, rec.Code)
	assert.Equal(t, "Method not allowed\n", rec.Body.String())
}

func TestSignupPreflightPostMethod(t *testing.T) {
	req := httptest.NewRequest("POST", "/signup", nil)
	rec := httptest.NewRecorder()

	handler := http.HandlerFunc(SignupHandler)
	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestSigninPreflightPostMethod(t *testing.T) {
	req := httptest.NewRequest("POST", "/signin", nil)
	rec := httptest.NewRecorder()

	handler := http.HandlerFunc(SigninHandler)
	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestSignupPreflightOptionsMethod(t *testing.T) {
	req := httptest.NewRequest("OPTIONS", "/signup", nil)
	rec := httptest.NewRecorder()

	handler := http.HandlerFunc(SignupHandler)
	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "success", rec.Body.String())
}

func TestSigninPreflightOptionsMethod(t *testing.T) {
	req := httptest.NewRequest("OPTIONS", "/signin", nil)
	rec := httptest.NewRecorder()

	handler := http.HandlerFunc(SigninHandler)
	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "success", rec.Body.String())
}
