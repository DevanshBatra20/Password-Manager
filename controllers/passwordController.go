package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/DevanshBatra20/Password-Manager/configs"
	"github.com/DevanshBatra20/Password-Manager/helpers"
	"github.com/DevanshBatra20/Password-Manager/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var passwordCollection *mongo.Collection = configs.GetCollection(configs.DB, "user_passwords")
var user *mongo.Collection = configs.GetCollection(configs.DB, "user_details")

func CreatePassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		userId := c.Param("user_id")
		objectId, err := primitive.ObjectIDFromHex(userId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Invalid userId",
				"data":    err.Error(),
			})
			return
		}

		var result bson.M

		err = user.FindOne(ctx, bson.M{"_id": objectId}).Decode(&result)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "User with " + userId + " not found",
				"data":    err.Error(),
			})
			return
		}

		var password models.Password
		if err := c.BindJSON(&password); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "error",
				"data":    err.Error(),
			})
			return
		}

		encryptedPassword, err := helpers.EncryptAES(password.Password_Value)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Some problem occured",
				"data":    err.Error(),
			})
			return
		}

		validateErr := validate.Struct(&password)
		if validateErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "error",
				"data":    validateErr.Error(),
			})
			return
		}

		newPassword := models.Password{
			ID:             primitive.NewObjectID(),
			Password_Name:  password.Password_Name,
			Password_Value: encryptedPassword,
			Password_Type:  password.Password_Type,
			UserId:         objectId,
		}

		_, err = passwordCollection.InsertOne(ctx, newPassword)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Some problem occured",
				"data":    err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"status":  http.StatusCreated,
			"message": "Password added successfully",
			"data":    newPassword,
		})
	}
}
