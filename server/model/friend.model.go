package model

import (
	"time"
)

type Friend struct {
	UserID    uint `gorm:"primaryKey"`
	FriendID  uint `gorm:"primaryKey"`
	Status    uint `gorm:"type:tinyint"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
