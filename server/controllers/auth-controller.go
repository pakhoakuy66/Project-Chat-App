package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"server/models"
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
	user, exists := ctx.Get("newuser")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "something went wrong"})
		return
	}
	result := models.DB.Create(user.(*models.User))
	if err := result.Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": user.(*models.User)})
}

func Login(ctx *gin.Context) {
	accessToken, exists := ctx.Get("accesstoken")
	refreshToken, exists2 := ctx.Get("refreshtoken")
	maxage, exists3 := ctx.Get("maxage")
	if !exists || !exists2 || !exists3 {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "something went wrong"})
		return
	}
	ctx.SetCookie(
		"AccessToken",
		accessToken.(string),
		maxage.(int),
		"/",
		"localhost",
		false,
		true,
	)
	ctx.SetCookie(
		"RefreshToken",
		refreshToken.(string),
		maxage.(int),
		"/",
		"localhost",
		false,
		true,
	)
	ctx.JSON(http.StatusOK, gin.H{
		"accesstoken":  accessToken,
		"refreshtoken": refreshToken,
	})
}

func Logout(ctx *gin.Context) {
	ctx.SetCookie(
		"AccessToken",
		"",
		-1,
		"/",
		"localhost",
		false,
		true,
	)
	ctx.SetCookie(
		"RefreshToken",
		"",
		-1,
		"/",
		"localhost",
		false,
		true,
	)
}
