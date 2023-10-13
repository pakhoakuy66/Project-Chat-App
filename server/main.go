package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"server/models"
)

var port, username, password, database string

func main() {
	setupEnvironment()
	models.ConnectDatabase(username, password, database)
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"Hello": "World",
		})
	})
	router.Run(fmt.Sprintf("localhost:%s", port))
}

func setupEnvironment() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	} else {
		fmt.Println("Successfully loaded .env file")
	}
	port = os.Getenv("PORT")
	username = os.Getenv("DBUSERNAME")
	password = os.Getenv("DBPASSWORD")
	database = os.Getenv("DATABASE")
}
