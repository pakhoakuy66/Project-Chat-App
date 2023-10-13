package models

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(username string, password string, dbname string) {
	dsn := fmt.Sprintf("%[1]s:%[2]s@tcp(127.0.0.1:3306)/%[3]s?charset=utf8mb4&parseTime=True&loc=Local", username, password, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database")
	} else {
		fmt.Println("Successfully connect to the database")
	}
	DB = db
}
