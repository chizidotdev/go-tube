package main

import (
	"fmt"
	"log"

	"github.com/chizidotdev/go-tube/controllers"
	"github.com/chizidotdev/go-tube/database"
	"github.com/chizidotdev/go-tube/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	database.Connect()
}

func main() {
	// ...
	fmt.Println("Hello, GoTube!")
	// ...

	r := gin.Default()
	r.POST("/api/v1/signup", controllers.Signup)
	r.POST("/api/v1/signin", controllers.Login)
	r.GET("/api/v1/validatetoken", middlewares.AuthMiddleware, controllers.ValidateToken)

	r.Run()
}
