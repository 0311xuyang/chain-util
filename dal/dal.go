package dal

import (
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	// DB is the database connection.
	DB   *gorm.DB
	once sync.Once
)

func init() {
	once.Do(func() {
		DB = ConnectDB()
	})
}

// ConnectDB connects to the database.
func ConnectDB() *gorm.DB {
	dsn := "root:123456@tcp(127.0.0.1:3306)/toolkit?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	Prepare(db)
	return db
}
