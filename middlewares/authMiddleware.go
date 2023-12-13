package middlewares

import (
	"net/http"
	"strings"

	"github.com/DevanshBatra20-PasswordManager/helpers"
)

func AuthMiddleware() func(handler http.Handler) http.Handler {
	return func(inner http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

			token := extractToken(request)

			_, err := helpers.ValidateToken(token)
			if err != nil {

				http.Error(writer, err.Error(), http.StatusUnauthorized)
				return
			}

			inner.ServeHTTP(writer, request)
		})
	}
}

func extractToken(request *http.Request) string {
	authHeader := request.Header.Get("Authorization")

	if authHeader != "" {
		parts := strings.Split(authHeader, " ")
		if len(parts) == 2 && parts[0] == "Bearer" {
			return parts[1]
		}
	}

	return ""
}
