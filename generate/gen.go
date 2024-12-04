package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	dsn := "root:123456@tcp(127.0.0.1:3306)/toolkit?charset=utf8mb4&parseTime=True"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	prepare(db)
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:       "./dal/dao",
		Mode:          gen.WithoutContext,
		FieldSignable: true,
	})
	// generate code
	g.UseDB(db)
	g.GenerateAllTable()
	g.ApplyBasic(g.GenerateAllTable()...)
	g.Execute()
}
