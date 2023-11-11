package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"server/model"
	"server/route"
	"server/service"
)

var port, username, password, database string

var corsConfig cors.Config

func main() {
	setupEnvironment()
	model.ConnectDatabase(username, password, database)
	r := gin.Default()
	r.Use(cors.New(corsConfig))
	route.InitAuthRoute(r)
	route.InitFriendsRoute(r)
	r.Run(":" + port)
}

func setupEnvironment() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(".env not found")
	}
	port = os.Getenv("PORT")
	username = os.Getenv("DBUSERNAME")
	password = os.Getenv("DBPASSWORD")
	database = os.Getenv("DATABASE")
	service.SetJwtKey(os.Getenv("JWTKEY"))
	corsConfig = cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization")
}
