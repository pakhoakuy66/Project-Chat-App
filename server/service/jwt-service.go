package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"server/model"
)

type Claims struct {
	ID          uint      `json:"id"`
	Username    string    `json:"username"`
	Gender      string    `json:"gender"`
	FirstName   string    `json:"firstname"`
	LastName    string    `json:"lastname"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phonenumber"`
	BirthDay    time.Time `json:"birthday"`
	CreatedAt   time.Time `json:"createdat"`
	jwt.RegisteredClaims
}

var jwtKey []byte

func SetJwtKey(key string) {
	jwtKey = []byte(key)
}

func GenerateToken(user *model.User, expirationTime time.Time) (string, error) {
	claims := Claims{
		ID:          user.ID,
		Username:    user.Username,
		Gender:      user.GenderStr(),
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		BirthDay:    user.BirthDay,
		CreatedAt:   user.CreatedAt,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	return token.SignedString(jwtKey)
}
