package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestRegister tests the /api/register endpoint
func TestRegister(t *testing.T) {
	// Create a request body with user registration data
	requestBody := []byte(`{"username": "testuser", "password": "password123"}`)

	// Create a new HTTP request with the request body
	req, err := http.NewRequest("POST", "/api/register", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()

	// Create a mock HTTP handler using the Register handler function
	handler := http.HandlerFunc(Register)

	// Serve the HTTP request to the mock HTTP handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	var response map[string]string
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf("error decoding response body: %v", err)
	}

	// Check if the response contains the expected message
	expected := "User registered successfully"
	if response["message"] != expected {
		t.Errorf("handler returned unexpected message: got %v want %v", response["message"], expected)
	}
}

// TestLogin tests the /api/login endpoint
func TestLogin(t *testing.T) {
	// Create a request body with user login data
	requestBody := []byte(`{"username": "testuser", "password": "password123"}`)

	// Create a new HTTP request with the request body
	req, err := http.NewRequest("POST", "/api/login", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()

	// Create a mock HTTP handler using the Login handler function
	handler := http.HandlerFunc(Login)

	// Serve the HTTP request to the mock HTTP handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	var response map[string]string
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf("error decoding response body: %v", err)
	}

	// Check if the response contains the expected message
	expected := "User logged in successfully"
	if response["message"] != expected {
		t.Errorf("handler returned unexpected message: got %v want %v", response["message"], expected)
	}
}
