package controller

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"server/model"
	"server/service"
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

type Authorization struct {
	AccessToken  string `json:"accesstoken"`
	RefreshToken string `json:"refreshtoken"`
}

var accessTokenDuration = 10 * time.Minute

var refreshTokenDuration = accessTokenDuration * 2

func Register(ctx *gin.Context) {
	var req RegisterRequest
	var err error
	var hashedPassword string
	var user model.User
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
	hashedPassword, err = service.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	user = model.User{
		Username:    req.Username,
		Password:    hashedPassword,
		Gender:      req.Gender == "male",
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		BirthDay:    req.BirthDay,
	}
	result := model.DB.Create(&user)
	if err = result.Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusOK)
}

func Login(ctx *gin.Context) {
	var req LoginRequest
	var err error
	var result model.User
	var accessToken string
	var refreshToken string
	if err = ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err = model.DB.Where(&model.User{Username: req.Username}).First(&result).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	if err = service.CheckPassword(req.Password, result.Password); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	accessToken, err = service.GenerateToken(&result, time.Now().Add(accessTokenDuration))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	refreshToken, err = service.GenerateToken(&result, time.Now().Add(refreshTokenDuration))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, Authorization{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}
