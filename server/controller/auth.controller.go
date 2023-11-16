package controller

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"server/model"
	"server/service"
)

type registerRequest struct {
	Username    string    `json:"username" binding:"required"`
	Password    string    `json:"password" binding:"required"`
	Gender      string    `json:"gender" binding:"required"`
	FirstName   string    `json:"firstname" binding:"required"`
	LastName    string    `json:"lastname" binding:"required"`
	Email       string    `json:"email" binding:"required"`
	PhoneNumber string    `json:"phonenumber" binding:"required"`
	BirthDay    time.Time `json:"birthday" binding:"required"`
}

type loginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

var accessTokenDuration = 20 * time.Minute

func Register(ctx *gin.Context) {
	var req = registerRequest{}
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
	hashedPassword, err := service.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := model.DB.Create(&model.User{
		Username:    req.Username,
		Password:    hashedPassword,
		Gender:      req.Gender == "male",
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		BirthDay:    req.BirthDay,
	}).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusOK)
}

func Login(ctx *gin.Context) {
	var req = loginRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := model.User{}
	if err := model.DB.Where(&model.User{Username: req.Username}).First(&result).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	if err := service.CheckPassword(req.Password, result.Password); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	expiredTime := time.Now().Add(accessTokenDuration)
	creds, err := service.UserToCreds(&result, expiredTime)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, creds)
}

func Refresh(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	header := strings.Split(authHeader, " ")
	if len(header) != 2 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "incorrect argument in the Authorization header",
		})
		return
	}
	if header[0] != "Bearer" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "unexpected argument in the Authorization header",
		})
		return
	}
	tokenString := header[1]
	claims, err := service.TokenToClaims(tokenString)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	expiredTime := time.Now().Add(accessTokenDuration)
	newCreds, err := service.ClaimsToCreds(&claims, expiredTime)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, newCreds)
}

func Profile(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	header := strings.Split(authHeader, " ")
	if len(header) != 2 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "incorrect argument in the Authorization header",
		})
		return
	}
	if header[0] != "Bearer" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "unexpected argument in the Authorization header",
		})
		return
	}
	tokenString := header[1]
	claims, err := service.TokenToClaims(tokenString)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, claims)
}
