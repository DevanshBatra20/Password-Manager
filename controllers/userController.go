package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/DevanshBatra20/Password-Manager/configs"
	"github.com/DevanshBatra20/Password-Manager/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection = configs.GetCollection(configs.DB, "user_details")
var bucketName string = configs.EnvAwsBucketName()
var bucketRegion string = configs.EnvAwsBucketRegion()

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		userId := c.Param("user_id")
		var user models.User

		objId, _ := primitive.ObjectIDFromHex(userId)

		err := collection.FindOne(ctx, bson.M{"user_id": objId}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "error",
				"data":    err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "success",
			"data":    user,
		})
	}
}

func UploadImage() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		userId := c.Param("user_id")
		objId, _ := primitive.ObjectIDFromHex(userId)

		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": err.Error(),
			})
			return
		}

		filContent, err := file.Open()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": err.Error(),
			})
			return
		}
		defer filContent.Close()

		creds := credentials.NewEnvCredentials()
		_, err = creds.Get()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": err.Error(),
			})
			return
		}

		sess := session.Must(session.NewSession(
			&aws.Config{
				Credentials: creds,
				Region:      aws.String(bucketRegion),
			},
		))

		svc := s3.New(sess)

		_, err = svc.PutObject(&s3.PutObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String("User " + userId),
			Body:   filContent,
		})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": err.Error(),
			})
			return
		}

		imageUrl := "https://" + bucketName + ".s3." + bucketRegion + ".amazonaws.com/" + file.Filename

		_, err = collection.UpdateOne(ctx,
			bson.M{"user_id": objId},
			bson.M{"$set": bson.M{"image_url": imageUrl}})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Profile picture updated successfully!",
			"data":    imageUrl,
		})

	}
}
