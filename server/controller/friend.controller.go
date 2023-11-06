package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"server/model"
	"server/service"
)

func GetAllFriend(ctx *gin.Context) {
	userClaims, exists := ctx.Get("claims")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong",
		})
		return
	}
	claims := userClaims.(*service.Claims)
	var friends []model.User
	result := model.DB.
		Joins("JOIN friends ON friends.friend_id = users.id AND friends.status = 0").
		Where("friends.user_id = ?", claims.ID).
		Find(&friends)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"friends": friends,
	})
}
