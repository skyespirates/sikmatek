package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Id         int    `json:"id"`
	Email      string `json:"email"`
	RoleId     int    `json:"role_id"`
	ConsumerId int    `json:"consumer_id"`
	IsVerified bool   `json:"is_verified"`
	jwt.RegisteredClaims
}

type JwtPayload struct {
	Id         int
	Email      string
	RoleId     int
	ConsumerId int
	IsVerified bool
}

func GenerateToken(payload JwtPayload) string {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		Id:         payload.Id,
		Email:      payload.Email,
		RoleId:     payload.RoleId,
		ConsumerId: payload.ConsumerId,
		IsVerified: payload.IsVerified,
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
