package models

import (
	"time"
)

type User struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Username    string `gorm:"unique"`
	Password    string
	FirstName   string
	LastName    string
	Email       string `gorm:"unique,check:email LIKE '%@%'"`
	PhoneNumber string `gorm:"unique;check:(phone_number LIKE '0%' OR phone_number LIKE '+%') AND length(phone_number) > 6 AND length(phone_number) < 16 AND SUBSTRING(phone_number, 2) REGEXP '^[0-9]+$'"`
	BirthDay    time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
