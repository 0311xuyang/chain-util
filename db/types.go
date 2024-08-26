package db

// User is a struct that holds user information.
type User struct {
	ID    string `dynamo:"ID,hash"`
	Name  string `dynamo:"Name"`
	Email string `dynamo:"Email"`
}
