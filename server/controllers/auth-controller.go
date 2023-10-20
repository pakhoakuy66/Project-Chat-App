package controllers

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"server/models"
	"server/services"
)

type RegisterRequest struct {
	Username    string    `json:"username" binding:"required"`
	Password    string    `json:"password" binding:"required"`
	Gender      string    `json:"gender" binding:"required"`
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

func Register(ctx *gin.Context) {
	var req RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if strings.Contains(req.Username, " ") || strings.Contains(req.Password, " ") {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Username and Password must not contain spaces",
		})
		return
	}
	if req.Gender != "male" && req.Gender != "female" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Gender must either be male or female",
		})
		return
	}
	if req.Gender != "male" && req.Gender != "female" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Gender must either be male or female",
		})
		return
	}
	hashedPassword, err := services.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	user := models.User{
		UUID:        uuid.NewString(),
		Username:    req.Username,
		Password:    hashedPassword,
		Gender:      req.Gender == "male",
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
	tokenString, err := services.GenerateToken(&result)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.SetCookie(
		"Authorization",
		tokenString,
		int(services.ExpirationDuration()),
		"/",
		"localhost",
		false,
		true,
	)
	ctx.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func Refresh(ctx *gin.Context) {
	claims := ctx.MustGet("userclaims").(*services.UserClaims)
	tokenString, err := services.RefreshToken(claims)
	if err != nil {
		switch err.(type) {
		case *services.EarlyRefreshError:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	ctx.SetCookie(
		"Authorization",
		tokenString,
		int(services.ExpirationDuration()),
		"/",
		"localhost",
		false,
		true,
	)
}

func Logout(ctx *gin.Context) {
	ctx.SetCookie(
		"Authorization",
		"",
		-1,
		"/",
		"localhost",
		false,
		true,
	)
}
