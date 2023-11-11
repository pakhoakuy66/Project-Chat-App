package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"server/model"
)

type Claims struct {
	ID                   uint      `json:"id"`
	Username             string    `json:"username"`
	Gender               string    `json:"gender"`
	FirstName            string    `json:"firstname"`
	LastName             string    `json:"lastname"`
	Email                string    `json:"email"`
	PhoneNumber          string    `json:"phonenumber"`
	BirthDay             time.Time `json:"birthday"`
	CreatedAt            time.Time `json:"createdAt"`
	jwt.RegisteredClaims `json:"-"`
}

type Credentials struct {
	Jwt       string `json:"jwt"`
	ExpiredAt int64  `json:"expiredAt"`
}

var jwtKey []byte

func SetJwtKey(key string) {
	jwtKey = []byte(key)
}

func UserToCreds(user *model.User, expirationTime time.Time) (Credentials, error) {
	return ClaimsToCreds(
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
		},
		expirationTime,
	)
}

func ClaimsToCreds(claims *Claims, expirationTime time.Time) (Credentials, error) {
	claims.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtKey)
	return Credentials{Jwt: tokenStr, ExpiredAt: expirationTime.UnixMilli()}, err
}

func TokenToClaims(tokenStr string) (Claims, error) {
	var claims Claims
	_, err := jwt.ParseWithClaims(tokenStr, &claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return Claims{}, err
	}
	return claims, nil
}
