package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"server/models"
	"server/routes"
	"server/services"
)

var port, username, password, database string

var corsConfig cors.Config

func main() {
	setupEnvironment()
	models.ConnectDatabase(username, password, database)
	r := gin.Default()
	r.Use(cors.New(corsConfig))
	routes.InitAuthRoute(r)
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
	services.SetJwtKey(os.Getenv("JWTKEY"))
	corsConfig = cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{
        "http://localhost:5173",
        "http://localhost:80",
    }
	corsConfig.AllowCredentials = true
}
