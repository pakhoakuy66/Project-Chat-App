package middlewares

import "github.com/gin-gonic/gin"

func AllowCredentials(ctx *gin.Context) {
	ctx.Writer.Header().Add("Access-Control-Allow-Credentials", "true")
	ctx.Next()
}
