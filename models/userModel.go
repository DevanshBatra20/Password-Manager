package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Username      string             `json:"username" validate:"required"`
	First_Name    string             `json:"first_name" validate:"required"`
	Last_Name     string             `json:"last_name" validate:"required"`
	Phone         string             `json:"phone" validate:"required"`
	Email         string             `json:"email" validate:"required"`
	Password      string             `json:"password" validate:"required"`
	Image_Url     string             `json:"image_url"`
	Token         string             `json:"token"`
	Refresh_Token string             `json:"refresh_token"`
}
