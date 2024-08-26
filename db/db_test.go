package db_test

import (
	"testing"

	"github.com/0311xuyang/chain-util/db"
)

// TestCheckValid tests the CheckValid function.
func TestCheckValid(t *testing.T) {
	if !db.CheckValid("1") {
		t.Error("1 should be valid.")
	}
	if db.CheckValid("2") {
		t.Error("2 should not be valid.")
	}
}

func TestGetSetUser(t *testing.T) {
	// Create a new DB object.
	Store := db.New()

	// Create a new User object.
	user1 := &db.User{
		ID:    "1",
		Name:  "John Doe",
		Email: "user1@test.com",
	}
	err := Store.PutUser(user1)
	if err != nil {
		t.Error("Error putting user:", err)
	}
	user2, err := Store.GetUser("1")
	if err != nil {
		t.Error("Error getting user:", err)
	}
	if user1.ID != user2.ID || user1.Name != user2.Name || user1.Email != user2.Email {
		t.Error("User data does not match.")
	}
}
