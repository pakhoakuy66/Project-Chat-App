package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"server/controllers"
	"server/models"
	"server/routes"
)

var port, username, password, database string

func main() {
	setupEnvironment()
	models.ConnectDatabase(username, password, database)
	r := gin.Default()
	routes.InitUserRoute(r)
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
	controllers.SetJwtKey(os.Getenv("JWTKEY"))
}
