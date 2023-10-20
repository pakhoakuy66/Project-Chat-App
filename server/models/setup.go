package models

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(username string, password string, dbname string) {
	dsn := fmt.Sprintf("%[1]s:%[2]s@tcp(127.0.0.1:3306)/%[3]s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, dbname)
	var err error
	if DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		log.Fatal("Failed to connect to the database")
	}
	if err = DB.AutoMigrate(&User{}); err != nil {
		log.Fatal("Failed to migrate tables to the database")
	}
}
