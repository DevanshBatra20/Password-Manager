package helpers

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/DevanshBatra20-PasswordManager/exception"
	"github.com/golang-jwt/jwt/v4"
)

type SignedDetails struct {
	Email      string
	First_Name string
	Last_Name  string
	jwt.RegisteredClaims
}

var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenerateJwtToken(email string, firstName string, lastName string) (signedToken string, err error) {
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
		return "", err
	}

	return token, nil
}

func ExtractToken(request *http.Request) string {
	authHeader := request.Header.Get("Authorization")

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
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		return nil, exception.InvalidToken{Token: signedToken}
	}

	if claims.ExpiresAt.Time.Unix() < time.Now().Local().Unix() {
		return nil, exception.TokenExpired{Token: signedToken}
	}

	return claims, nil
}
