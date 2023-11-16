package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"server/service"
)

func Authorize(ctx *gin.Context) {
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
	ctx.Set("claims", &claims)
	ctx.Next()
}
