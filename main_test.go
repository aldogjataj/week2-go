package main

import (
	"acme/api"
	"acme/model"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestRootHandler(t *testing.T) {
	//Arrange
	req, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Fatal(err)
	}

	testResponse := httptest.NewRecorder()
	handler := http.HandlerFunc(rootHandler)

	// Act
	handler.ServeHTTP(testResponse, req)

	//Assert
	if status := testResponse.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Hello Airah"
	if testResponse.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", testResponse.Body.String(), expected)
	}
}

func TestGetUsersHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	testResponse := httptest.NewRecorder()
	handler := http.HandlerFunc(api.GetUsers)

	expected := []model.User{
		{ID: 1, Name: "User 1"},
		{ID: 2, Name: "User 2"},
		{ID: 3, Name: "User 3"},
	}
	if err != nil {
		t.Fatalf("Failed to marshal expected JSON: %v", err)
	}

	// Act
	handler.ServeHTTP(testResponse, req)

	// Assert
	if status := testResponse.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var actual []model.User
	if err := json.Unmarshal(testResponse.Body.Bytes(), &actual); err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func TestRootHandlerWithServer(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(rootHandler))
	defer server.Close()

	response, err := http.Get(server.URL + "/")
	if err != nil {
		t.Fatalf("Failed to send GET request: %v", err)
	}

	defer response.Body.Close()

	if status := response.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Hello Airah"
	bodyBytes, err := io.ReadAll(response.Body)

	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}
	if string(bodyBytes) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", string(bodyBytes), expected)
	}
}

// integration tests
func TestGetUsersHandlerWithServer(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(api.GetUsers))
	defer server.Close()

	response, err := http.Get(server.URL + "/api/users")
	if err != nil {
		t.Fatalf("Failed to send GET request: %v", err)
	}

	defer response.Body.Close()

	if status := response.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	body, errRead := io.ReadAll(response.Body)
	if errRead != nil {
		t.Fatalf("Error reading response body: %v", err)
	}

	expected := []model.User{
		{ID: 1, Name: "User 1"},
		{ID: 2, Name: "User 2"},
		{ID: 3, Name: "User 3"},
	}
	var actual []model.User
	if err := json.Unmarshal(body, &actual); err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

// buggy
func TestCreateUserHandlerWithServer(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(api.CreateUser))
	defer server.Close()

	testReqBody := []byte(`{"name": "Test User"}`)

	response, err := http.Post(server.URL+"/api/users", "json", bytes.NewBuffer(testReqBody))
	if err != nil {
		t.Fatalf("Failed to send POST request: %v", err)
	}

	defer response.Body.Close()

	if status := response.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "User created successfully: 4"
	body, errRead := io.ReadAll(response.Body)
	if errRead != nil {
		t.Fatalf("Error reading response body: %v", err)
	}

	if string(body) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", string(body), expected)
	}
}
