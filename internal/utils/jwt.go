package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type JwtPayload struct {
	Id    int
	Email string
}

func GenerateToken(payload JwtPayload) string {
	expirationTime := time.Now().Add(24 * time.Hour) // token valid for 24h

	claims := &Claims{
		Id:    payload.Id,
		Email: payload.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		panic(err)
	}

	return tokenString
}

func VerifyToken(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (any, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}
