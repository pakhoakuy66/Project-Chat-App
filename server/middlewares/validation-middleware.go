package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"server/controllers"
	"server/models"
	"server/services"
)

func Validate(ctx *gin.Context) {
	var req controllers.RegisterRequest
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
	ctx.Set("newuser", &user)
	ctx.Next()
}
