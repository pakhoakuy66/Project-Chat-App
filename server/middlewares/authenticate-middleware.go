package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"server/controllers"
	"server/models"
	"server/services"
)

func Authenticate(ctx *gin.Context) {
	var req controllers.LoginRequest
	var result models.User
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.DB.Where(&models.User{Username: req.Username}).First(&result).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	if err := services.CheckPassword(req.Password, result.Password); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	accessToken, err := services.GenerateToken(&result, time.Now().Add(10*time.Minute))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	refreshToken, err := services.GenerateToken(&result, time.Now().Add(20*time.Minute))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Set("accesstoken", accessToken)
	ctx.Set("refreshtoken", refreshToken)
	ctx.Set("maxage", int((20 * time.Minute).Seconds()))
	ctx.Next()
}
