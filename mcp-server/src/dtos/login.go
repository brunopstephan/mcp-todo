package dtos

import "github.com/golang-jwt/jwt/v5"

// Claims personalizados
type UserClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// Input/Output do login
type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginOutput struct {
	Token string `json:"token"`
}
