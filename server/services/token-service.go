package services

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"server/models"
)

type UserClaims struct {
	UUID        string    `json:"uuid"`
	Username    string    `json:"username"`
	Gender      string    `json:"gender"`
	FirstName   string    `json:"firstname"`
	LastName    string    `json:"lastname"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phonenumber"`
	BirthDay    time.Time `json:"birthday"`
	jwt.RegisteredClaims
}

type EarlyRefreshError struct{}

func (err *EarlyRefreshError) Error() string {
	return "Too early to refresh this token"
}

var jwtKey []byte

func JwtKey() []byte {
	return jwtKey
}

func SetJwtKey(key string) {
	jwtKey = []byte(jwtKey)
}

var expirationDuration time.Duration = 5 * time.Minute

func ExpirationDuration() time.Duration {
	return expirationDuration
}

var renewAt time.Duration = 30 * time.Second

func GenerateToken(user *models.User) (tokenString string, err error) {
	expirationTime := time.Now().Add(expirationDuration)
	claims := UserClaims{
		UUID:        user.UUID,
		Username:    user.Username,
		Gender:      user.GenderStr(),
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		BirthDay:    user.BirthDay,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	return token.SignedString(jwtKey)
}

func RefreshToken(claims *UserClaims) (string, error) {
	if time.Until(claims.ExpiresAt.Time) > renewAt {
		return "", &EarlyRefreshError{}
	}
	expirationTime := time.Now().Add(expirationDuration)
	claims.ExpiresAt = jwt.NewNumericDate(expirationTime)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
