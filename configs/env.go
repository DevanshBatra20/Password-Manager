package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMomgoUri() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("MONGO_URI")
}

func EnvSecretKey() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("SECRET_KEY")
}

func EnvAwsBucketName() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading the .env file")
	}

	return os.Getenv("AWS_BUCKET_NAME")
}

func EnvAwsBucketRegion() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading the .env file")
	}

	return os.Getenv("AWS_BUCKET_REGION")
}

func EnvAesSecretKey() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading the .env file")
	}

	return os.Getenv("AES_SECRET_KEY")
}
