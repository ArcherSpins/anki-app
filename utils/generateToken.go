package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

type Claims struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateJWT(id int, username string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	tokenKey := os.Getenv("SECRET_KEY")
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		ID:       id,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println(tokenKey, token)

	tokenString, err := token.SignedString([]byte(tokenKey))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJWT(tokenString string) (*Claims, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	claims := &Claims{}
	tokenKey := os.Getenv("SECRET_KEY")

	if tokenKey == "" {
		fmt.Println("SECRET_KEY is not set")
		return &Claims{}, errors.New("SECRET_KEY is not set")
	}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}
