package model

import (
	"time"
)

type User struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Username    string `gorm:"type:varchar(50);unique;index;check username NOT LIKE '% %'"`
	Password    string `gorm:"type:varchar(60);check password NOT LIKE '% %'"`
	Gender      bool
	FirstName   string    `gorm:"type:nvarchar(50)"`
	LastName    string    `gorm:"type:nvarchar(50)"`
	Email       string    `gorm:"type:varchar(100);unique,check:email LIKE '%@%'"`
	PhoneNumber string    `gorm:"type:varchar(15);unique;check:phone_number REGEXP '^[0+][0-9]{6,}$'"`
	BirthDay    time.Time `gorm:"type:date"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Friends     []Friend `gorm:"foreignKey:UserID;references:ID;foreignKey:FriendID;references:ID"`
}

func (user *User) GenderStr() string {
	if user.Gender {
		return "male"
	}
	return "female"
}
