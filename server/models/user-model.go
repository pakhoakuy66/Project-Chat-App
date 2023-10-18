package models

import (
	"time"
)

type User struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Username    string `gorm:"type:varchar(50);unique"`
	Password    string `gorm:"type:char(60)"`
	Gender      bool
	FirstName   string    `gorm:"type:nvarchar(50)"`
	LastName    string    `gorm:"type:nvarchar(50)"`
	Email       string    `gorm:"type:nvarchar(100);unique,check:email LIKE '%@%'"`
	PhoneNumber string    `gorm:"type:varchar(15);unique;check:phone_number REGEXP '^[0+][0-9]{6,}$'"`
	BirthDay    time.Time `gorm:"type:date"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
