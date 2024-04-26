package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/DevanshBatra20/Password-Manager/configs"
	"github.com/DevanshBatra20/Password-Manager/helpers"
	"github.com/DevanshBatra20/Password-Manager/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "user_details")
var validate = validator.New()

func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "error",
				"data":    err.Error(),
			})
			return
		}

		validateErr := validate.Struct(&user)
		if validateErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "error",
				"data":    validateErr.Error(),
			})
			return
		}

		countEmail, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "error",
				"data":    err.Error(),
			})
		}

		if countEmail > 0 {
			c.JSON(http.StatusForbidden, gin.H{
				"status":  http.StatusForbidden,
				"message": "error",
				"data":    "User with the same email already exists",
			})
			return
		}

		countPhone, err := userCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "error",
				"data":    err.Error(),
			})
		}

		if countPhone > 0 {
			c.JSON(http.StatusForbidden, gin.H{
				"status":  http.StatusForbidden,
				"message": "error",
				"data":    "User with same phone number already exists",
			})
		}

		countUsername, err := userCollection.CountDocuments(ctx, bson.M{"username": user.Username})
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "error",
				"data":    err.Error(),
			})
		}

		if countUsername > 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": "error",
				"data":    "User with same username already exists",
			})
		}

		signedToken, refreshToken, err := helpers.GenerateJwtToken(user.First_Name, user.First_Name, user.Last_Name)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "error",
				"data":    err.Error(),
			})
			return
		}

		hashPassword, err := helpers.HashPassword(user.Password)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "error",
				"data":    err.Error(),
			})
			return
		}

		newUser := models.User{
			ID:            primitive.NewObjectID(),
			First_Name:    user.First_Name,
			Last_Name:     user.Last_Name,
			Username:      user.Username,
			Phone:         user.Phone,
			Email:         user.Email,
			Password:      hashPassword,
			Token:         signedToken,
			Refresh_Token: refreshToken,
		}

		result, err := userCollection.InsertOne(ctx, newUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "error",
				"data":    err.Error(),
			})
			return
		}

		err = userCollection.FindOne(ctx, bson.M{"_id": result.InsertedID}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "error",
				"data":    err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"status":  http.StatusCreated,
			"message": "User created successfully!",
			"data":    user,
		})
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var loginCredentials models.Login
		var user models.User
		if err := c.BindJSON(&loginCredentials); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "error",
				"data":    err.Error(),
			})
			return
		}

		validateErr := validate.Struct(&loginCredentials)
		if validateErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "error",
				"data":    validateErr.Error(),
			})
			return
		}

		err := userCollection.FindOne(ctx, bson.M{"email": loginCredentials.Email}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "error",
				"data":    err.Error(),
			})
			return
		}

		isPasswordValid := helpers.VerifyPassword(user.Password, loginCredentials.Password)
		if !isPasswordValid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "Incorrect Password",
				"data":    "Password you entered is incorrect",
			})
			return
		}

		token, refreshToken, err := helpers.GenerateJwtToken(user.Email, user.First_Name, user.Last_Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "error",
				"data":    err.Error(),
			})
			return
		}

		update := bson.M{
			"$set": bson.M{
				"token":         token,
				"refresh_token": refreshToken,
			},
		}

		_, err = userCollection.UpdateOne(ctx, bson.M{"_id": user.ID}, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "error",
				"data":    err.Error(),
			})
			return
		}

		user.Token = token
		user.Refresh_Token = refreshToken

		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Login successful",
			"data":    user,
		})

	}
}
