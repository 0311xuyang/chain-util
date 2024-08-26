package main

import (
	"fmt"

	"github.com/0311xuyang/chain-util/db"
)

func main() {
	// Create a new DB object.
	Store := db.New()

	// Create a new User object.
	user := &db.User{
		ID:    "1",
		Name:  "John Doe",
		Email: "user1@mail.com",
	}
	Store.PutUser(user)
	fmt.Println("User created.")

	// Get the user by ID.
	user, err := Store.GetUser("1")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("User:", user)
}
