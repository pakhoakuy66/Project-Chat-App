package middleware

import "github.com/gin-gonic/gin"

func SetAccessControlHeader(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173, http://localhost:80")
	ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
}
