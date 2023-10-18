package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"server/models"
	"server/services"
)

type RegisterRequest struct {
	Username    string    `json:"username" binding:"required"`
	Password    string    `json:"password" binding:"required"`
	Gender      bool      `json:"gender" binding:"required"`
	FirstName   string    `json:"firstname" binding:"required"`
	LastName    string    `json:"lastname" binding:"required"`
	Email       string    `json:"email" binding:"required"`
	PhoneNumber string    `json:"phonenumber" binding:"required"`
	BirthDay    time.Time `json:"birthday" binding:"required"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserClaims struct {
	ID          uint      `json:"id"`
	Username    string    `json:"username"`
	FirstName   string    `json:"firstname"`
	LastName    string    `json:"lastname"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phonenumber"`
	BirthDay    time.Time `json:"birthday"`
	jwt.RegisteredClaims
}

var jwtKey []byte

func SetJwtKey(key string) {
	jwtKey = []byte(key)
}

func Register(ctx *gin.Context) {
	var req RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashedPassword, err := services.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	user := models.User{
		Username:    req.Username,
		Password:    hashedPassword,
		Gender:      req.Gender,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		BirthDay:    req.BirthDay,
	}
	result := models.DB.Create(&user)
	if err := result.Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": req})
}

func Login(ctx *gin.Context) {
	var req LoginRequest
	var result models.User
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if err := models.DB.Where(&models.User{Username: req.Username}).First(&result).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	if err := services.CheckPassword(req.Password, result.Password); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := UserClaims{
		ID:          result.ID,
		Username:    result.Username,
		FirstName:   result.FirstName,
		LastName:    result.LastName,
		Email:       result.Email,
		PhoneNumber: result.PhoneNumber,
		BirthDay:    result.BirthDay,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	expirationDuration := expirationTime.Sub(time.Now())
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173, http://localhost:80")
	ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	ctx.SetCookie("Authorization", "Bearer "+tokenString,
		int(expirationDuration.Seconds()), "/", "localhost", false, true)
	ctx.Status(http.StatusOK)
}
