package utils

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func getJWTSecret() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return nil
	}
	return []byte(secret)
}

func GetTokenExpirationMinutes() int {
	expirationStr := os.Getenv("TOKEN_EXPIRATION_MINUTES")
	if expirationStr == "" {
		return 60
	}
	expiration, err := strconv.Atoi(expirationStr)
	if err != nil {
		return 60
	}
	return expiration
}

func GenerateToken(userID int, email string, expirationMinutes int) (string, error) {
	jwtSecret := getJWTSecret()
	if len(jwtSecret) == 0 {
		return "", errors.New("JWT_SECRET não configurado")
	}

	claims := Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expirationMinutes) * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "cortaai-auth-service",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", fmt.Errorf("erro ao assinar token: %v", err)
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (*Claims, error) {
	jwtSecret := getJWTSecret()
	if len(jwtSecret) == 0 {
		return nil, errors.New("JWT_SECRET não configurado")
	}

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, fmt.Errorf("erro ao parsear token: %v", err)
	}

	if !token.Valid {
		return nil, errors.New("token inválido")
	}

	return claims, nil
}
