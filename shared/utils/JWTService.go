package utils

import (
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
	"github.com/kritpi/499-senior-project-trip-service/property"
	"github.com/kritpi/499-senior-project-trip-service/shared/model"
)

func GenerateToken(googleClaims domain.GoogleUserClaims, now time.Time, cfg property.Property) (*string, error) {
	jwtSecrets := []byte(cfg.Auth.JWTSecrets)

	claims := model.JWTCustomClaims{
		ID:    googleClaims.Sub,
		Email: googleClaims.Email,
		Name:  googleClaims.Name,
		ImageUrl: googleClaims.ImageUrl,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{now.Add(time.Hour * 24)},
			IssuedAt:  &jwt.NumericDate{now},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecrets)
	if err != nil {
		log.Errorf("unable to generate JWT: %+v", err)
		return nil, err
	}
	return &tokenString, nil
}

func ParseAndValidateToken(tokenString string, cfg property.Property) (*model.JWTCustomClaims, error) {
	secret := []byte(cfg.Auth.JWTSecrets)

	token, err := jwt.ParseWithClaims(
		tokenString,
		&model.JWTCustomClaims{},
		func(t *jwt.Token) (interface{}, error) {
			if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
				return nil, jwt.ErrTokenSignatureInvalid
			}
			return secret, nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*model.JWTCustomClaims)
	if !ok || !token.Valid {
		return nil, jwt.ErrTokenSignatureInvalid
	}

	return claims, nil
}
