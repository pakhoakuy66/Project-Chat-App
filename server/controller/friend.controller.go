package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"server/model"
	"server/service"
)

type friendRelation struct {
	ID          uint      `json:"id"`
	Username    string    `json:"username"`
	Gender      string    `json:"gender"`
	FirstName   string    `json:"firstname"`
	LastName    string    `json:"lastname"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phonenumber"`
	BirthDay    time.Time `json:"birthday"`
	CreatedAt   time.Time `json:"createdAt"`
	Status      string    `json:"status"`
}

func GetAllRelationShip(ctx *gin.Context) {
	userClaims, exists := ctx.Get("claims")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong",
		})
		return
	}
	claims := userClaims.(*service.Claims)
	users := []friendRelation{}
	if err := model.DB.
		Model(&model.User{}).
		Select(`id, username, first_name, last_name, email, phone_number, birth_day, users.created_at,
            CASE WHEN gender = 0 THEN 'female' ELSE 'male' END AS gender,
            CASE
                WHEN friends.status = 0 THEN 'friend'
                WHEN friends.status = 1 THEN 'sent request'
                WHEN friends.status = 2 THEN 'recieved request'
                ELSE 'stranger'
            END AS 'status'`).
		Joins("LEFT JOIN friends ON friend_id = id AND user_id = ?", claims.ID).
		Find(&users, "id != ?", claims.ID).
		Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong",
			"error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func MakeFriendRequest(ctx *gin.Context) {
	userClaims, exists := ctx.Get("claims")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong",
		})
		return
	}
	claims := userClaims.(*service.Claims)
	friendId, err := strconv.Atoi(ctx.Param("reciever-id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := model.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&model.Friend{
			UserID:   claims.ID,
			FriendID: uint(friendId),
			Status:   1,
		}).Error; err != nil {
			return err
		}
		if err := tx.Create(&model.Friend{
			UserID:   uint(friendId),
			FriendID: claims.ID,
			Status:   2,
		}).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong",
			"error":   err.Error(),
		})
		return
	}
	ctx.Status(http.StatusOK)
}

func AcceptFriendRequest(ctx *gin.Context) {
	userClaims, exists := ctx.Get("claims")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong",
		})
		return
	}
	claims := userClaims.(*service.Claims)
	friendId, err := strconv.Atoi(ctx.Param("sender-id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := model.DB.
		Where(&model.Friend{
			UserID:   claims.ID,
			FriendID: uint(friendId),
			Status:   2,
		}).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "you are not allowed to accept this request",
			"error":   err.Error(),
		})
		return
	}
	if err := model.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.
			Model(&model.Friend{}).
			Where(&model.Friend{
				UserID:   claims.ID,
				FriendID: uint(friendId),
			}).
			Update("status", 0).
			Error; err != nil {
			return err
		}
		if err := tx.
			Model(&model.Friend{}).
			Where(&model.Friend{
				UserID:   uint(friendId),
				FriendID: claims.ID,
			}).
			Update("status", 0).
			Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "request does not exists",
			"error":   err.Error(),
		})
		return
	}
	ctx.Status(http.StatusOK)
}
