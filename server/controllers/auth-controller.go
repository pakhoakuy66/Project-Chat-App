package controllers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

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

type AuthorizationCookie struct {
	AccessToken  string `json:"accesstoken"`
	RefreshToken string `json:"refreshtoken"`
}

var accessTokenDuration = 10 * time.Minute

var refreshTokenDuration = accessTokenDuration * 2

func Register(ctx *gin.Context) {
	var req RegisterRequest
	var err error
	var hashedPassword string
	var user models.User
	if err = ctx.ShouldBindJSON(&req); err != nil {
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
	hashedPassword, err = services.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	user = models.User{
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
	if err = result.Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusOK)
}

func Login(ctx *gin.Context) {
	var req LoginRequest
	var err error
	var result models.User
	var accessToken string
	var refreshToken string
	var authorizationCookie []byte
	if err = ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err = models.DB.Where(&models.User{Username: req.Username}).First(&result).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	if err = services.CheckPassword(req.Password, result.Password); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	accessToken, err = services.GenerateToken(&result, time.Now().Add(accessTokenDuration))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	refreshToken, err = services.GenerateToken(&result, time.Now().Add(refreshTokenDuration))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	authorizationCookie, err = json.Marshal(AuthorizationCookie{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.SetCookie(
		"Authorization",
		string(authorizationCookie),
		int(refreshTokenDuration.Seconds()),
		"/",
		"localhost",
		false,
		true,
	)
	ctx.Status(http.StatusOK)
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
