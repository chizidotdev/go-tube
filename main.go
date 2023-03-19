package main

import (
	"fmt"

	"github.com/chizidotdev/go-tube/database"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func init() {
	client = database.Connect()
}

func main() {
	// ...
	fmt.Println("Hello, world!")
}
