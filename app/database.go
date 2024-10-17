package app

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func NewDB() *gorm.DB {
	dsn := os.Getenv("MYSQL")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Connection to database failed")
	}
	return db
}
