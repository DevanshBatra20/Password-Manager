package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Password struct {
	ID             primitive.ObjectID `bson:"_id"`
	Password_Name  string             `json:"password_name" validate:"required"`
	Password_Value string             `json:"password_value" validate:"required"`
	Password_Type  string             `json:"password_type" validate:"required"`
	UserId         primitive.ObjectID `bson:"user_id, omitempty"`
}
