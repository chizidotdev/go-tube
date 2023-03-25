package middlewares

import (
	"fmt"
	"net/http"

	"github.com/chizidotdev/go-tube/database"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
)

func AuthMiddleware(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	fmt.Println("token", tokenString)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, err
		}
		return []byte("secret"), nil
	})
	fmt.Println(token, err)

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userCollection := database.Database.Collection("users")
		err := userCollection.FindOne(c, bson.M{"email": claims["email"]})
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("user", claims)
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}
