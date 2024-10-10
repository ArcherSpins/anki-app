package middleware

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Извлечение токена из заголовка Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		// Ожидаемый формат заголовка: "Bearer <token>"
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		err := godotenv.Load()
		if err != nil {
			log.Println("No .env file found, relying on environment variables")
		}
		// Парсинг и проверка токена
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			secretKey := os.Getenv("SECRET_KEY")
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Продолжить выполнение запроса
		c.Next()
	}
}
