package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/create-user-test-case/models"
)

func TestCreateUser(t *testing.T) {
	user := models.User{
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Password: "password123",
	}
	payload, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(payload))
	rr := httptest.NewRecorder()

	// Handler function to test
	handler := http.HandlerFunc(handlers.CreateUser)
	handler.ServeHTTP(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	// Check the response body
	var createdUser models.User
	json.NewDecoder(rr.Body).Decode(&createdUser)
	if createdUser.Name != user.Name {
		t.Errorf("handler returned unexpected body: got %v want %v",
			createdUser.Name, user.Name)
	}
}
