package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID          primitive.ObjectID `bson:"_id"`
	Username    string             `bson:"username" validate:"required"`
	Email       string             `bson:"email" validate:"required,email,unique"`
	Password    string             `bson:"password, omitempty"`
	Image       string             `bson:"image"`
	Subscribers []string           `bson:"subscribers"`
}

type Video struct {
	ID          primitive.ObjectID `bson:"_id"`
	Author      string             `bson:"author" validate:"required"`
	Title       string             `bson:"title" validate:"required"`
	Description string             `bson:"description"`
	VideoURL    string             `bson:"video_url" validate:"required"`
	ImageURL    string             `bson:"image_url"`
	views       int                `bson:"views"`
	tags        []string           `bson:"tags"`
	likes       []string           `bson:"likes"`
	dislikes    []string           `bson:"dislikes"`
}

type Comment struct {
	ID      primitive.ObjectID `bson:"_id"`
	Author  string             `bson:"author" validate:"required"`
	VideoID string             `bson:"video_id" validate:"required"`
	Content string             `bson:"content" validate:"required"`
}
