package helpers

import (
	"strings"
	"time"

	"github.com/DevanshBatra20/Password-Manager/configs"
	"github.com/DevanshBatra20/Password-Manager/exceptions"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type SignedDetails struct {
	Email      string
	First_Name string
	Last_Name  string
	jwt.RegisteredClaims
}
type RefreshClaims struct {
	Email string
	jwt.RegisteredClaims
}

var SECRET_KEY string = configs.EnvSecretKey()

func GenerateJwtToken(email string, firstName string, lastName string) (string, string, error) {
	claims := &SignedDetails{
		Email:      email,
		First_Name: firstName,
		Last_Name:  lastName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Hour * 24)),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", "", err
	}

	refreshClaims := &RefreshClaims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Hour * 24 * 7)),
		},
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", "", err
	}

	return token, refreshToken, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func VerifyPassword(userPassword string, providedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(providedPassword))
	return err == nil
}

func ExtractToken(context *gin.Context) string {
	authHeader := context.GetHeader("Authorization")
	if authHeader != "" {
		parts := strings.Split(authHeader, " ")
		if len(parts) == 2 && parts[0] == "Bearer" {
			return parts[1]
		}
	}

	return ""
}

func ValidateToken(signedToken string) (claims *SignedDetails, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		return nil, exceptions.InvalidToken{Token: signedToken}
	}

	if claims.ExpiresAt.Time.Unix() < time.Now().Local().Unix() {
		return nil, exceptions.TokenExpired{Token: signedToken}
	}

	return claims, nil
}
