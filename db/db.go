package db

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

// DynamoDB is a global variable that holds the DynamoDB client.
var DynamoDB *dynamo.DB

// TableUser is the name of the Users table.
const TableUser = "Users"

func init() {
	DynamoDB = dynamo.New(session.Must(session.NewSession()), &aws.Config{
		Region:      aws.String("us-east-1"),
		Endpoint:    aws.String("http://127.0.0.1:8000"),
		Credentials: credentials.NewStaticCredentials("dummy", "dummy", ""),
	})
}

// DB is a struct that holds the DynamoDB tables.
type DB struct {
	Users dynamo.Table
}

// New creates a new DB object.
func New() *DB {
	CreateTable()
	fmt.Println("DB inited.")
	return &DB{
		Users: DynamoDB.Table(TableUser),
	}
}

// CreateTable creates the Users table.
func CreateTable() error {
	exists := false
	fmt.Println("Checking if table exists:", TableUser)
	tables, _ := DynamoDB.ListTables().All()
	for _, table := range tables {
		if table == TableUser {
			exists = true
			break
		}
	}
	if !exists {
		fmt.Println("Creating table:", TableUser)
		return DynamoDB.CreateTable(TableUser, &User{}).Run()
	}
	return nil
}

// GetUser gets a user by ID.
func (db *DB) GetUser(id string) (*User, error) {
	var user User
	err := db.Users.Get("ID", id).One(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// PutUser puts a user into the database.
func (db *DB) PutUser(user *User) error {
	return db.Users.Put(user).Run()
}

// CheckValid checks if the ID is valid.
func CheckValid(id string) bool {
	return strings.Contains(id, "1")
}
