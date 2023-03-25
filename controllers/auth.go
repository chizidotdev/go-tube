package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/chizidotdev/go-tube/database"
	"github.com/chizidotdev/go-tube/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func ValidateToken(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Valid token"})
}

func Signup(c *gin.Context) {
	var body models.User

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 14)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	userCollection := database.Database.Collection("users")
	user := models.User{
		Email:       body.Email,
		Password:    string(hash),
		Username:    body.Username,
		Image:       body.Image,
		Subscribers: []string{},
	}
	result, err := userCollection.InsertOne(c, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": result})
}

func Login(c *gin.Context) {
	var body models.User

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	var user models.User
	userCollection := database.Database.Collection("users")
	err := userCollection.FindOne(c, bson.M{"email": body.Email}).Decode(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect credentials"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect credentials"})
		return
	}
	userWithoutPassword := models.User{
		ID:          user.ID,
		Email:       user.Email,
		Username:    user.Username,
		Image:       user.Image,
		Subscribers: user.Subscribers,
	}

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"sub":         user.ID,
		"username":    user.Username,
		"email":       user.Email,
		"image":       user.Image,
		"subscribers": user.Subscribers,
	})

	fmt.Println(os.Getenv("SECRET"))
	tokenString, err := token.SignedString(os.Getenv("SECRET"))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"user": userWithoutPassword})
}

func Logout(c *gin.Context) {
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logged out"})
}
