package utils

import (
	"fmt"
	"mcp-server/src/dtos"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func ParseJWT(tokenString string) (dtos.UserClaims, error) {
	if !strings.HasPrefix(tokenString, "Bearer ") {
		return dtos.UserClaims{}, fmt.Errorf("invalid authorization")
	}
	tokenStr := strings.TrimPrefix(tokenString, "Bearer ")

	claims := &dtos.UserClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (any, error) {
		return os.Getenv("JWT_SECRET"), nil
	})
	if err != nil || !token.Valid {
		return dtos.UserClaims{}, fmt.Errorf("invalid token: %w", err)
	}

	return *claims, nil
}
