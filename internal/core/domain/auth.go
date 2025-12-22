package domain

import "github.com/golang-jwt/jwt/v5"

type GoogleIdToken struct {
	IDToken string `json:"id_token"`
}

type GoogleAuthResponse struct {
	AccessToken string `json:"access_token"`
}

type GoogleUserClaims struct {
	Sub      string `json:"sub"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

type JWTCustomClaims struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
	jwt.RegisteredClaims
}
