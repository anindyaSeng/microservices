package domain

import (
	"net/http"
	"testing"
)

func TestGetUserNoUserFound(t *testing.T) {

	// Initialize
	InitUserDomain()
	// Execution
	user, err := GetUser(0)

	//verification
	if user != nil {
		t.Error("We were not expecting a user with id 0")
	}
	if err == nil {
		t.Error("We were expecting an error when user id is 0")
	}
	if err.StatusCode != http.StatusNotFound {
		t.Error("We were expecting 404 when the user is not found")
	}
}

func TestGetUserIfFound(t *testing.T) {

	// Initialize
	InitUserDomain()
	// Execution
	user, err := GetUser(123)

	//verification
	if user == nil {
		t.Error("We were  expecting a user with id 123")
	}
	if err != nil {
		t.Error("We were expecting no error when user id is 123")
	}
}
