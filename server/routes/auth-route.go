package routes

import (
	"github.com/gin-gonic/gin"

	"server/controllers"
	"server/middleware"
)

func InitAuthRoute(r *gin.Engine) {
	authRoute := r.Group("/auth")
	authRoute.POST("/register", controllers.Register)
	authRoute.POST("/login", middleware.SetAccessControlHeader, controllers.Login)
	authRoute.POST("/refresh", middleware.Authorize, middleware.SetAccessControlHeader, controllers.Refresh)
	authRoute.POST("/logout", controllers.Logout)
	// authRoute.POST("/getuser", func(ctx *gin.Context) {
	// 	type GetUserRequest struct {
	// 		Token string `json:"token" binding:"required"`
	// 	}
	// 	var req GetUserRequest
	// 	var claims services.UserClaims
	// 	if err := ctx.ShouldBindJSON(&req); err != nil {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// 	token, err := jwt.ParseWithClaims(req.Token, &claims, func(t *jwt.Token) (interface{}, error) {
	// 		return services.JwtKey(), nil
	// 	})
	// 	if err != nil {
	// 		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// 	if !token.Valid {
	// 		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "token is outdated"})
	// 		return
	// 	}
	// 	ctx.JSON(http.StatusOK, gin.H{"user": claims})
	// })
}
