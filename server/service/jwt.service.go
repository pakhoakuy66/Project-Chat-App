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

func GenerateTokenWithUser(user *model.User, expirationTime time.Time) (string, error) {
	return GenerateTokenWithClaims(
		&Claims{
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
		},
		expirationTime,
	)
}

func GenerateTokenWithClaims(claims *Claims, expirationTime time.Time) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func TokenToClaims(tokenStr string) (*Claims, error) {
	var claims Claims
	_, err := jwt.ParseWithClaims(tokenStr, &claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	return &claims, nil
}
