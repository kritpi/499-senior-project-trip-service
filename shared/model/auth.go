package model

import "github.com/golang-jwt/jwt/v5"

type JWTCustomClaims struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
	jwt.RegisteredClaims
}
