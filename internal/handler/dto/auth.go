package dto

import (
	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
	"github.com/kritpi/499-senior-project-trip-service/shared/model"
)

type JWTCustomClaims model.JWTCustomClaims

type GoogleIdToken struct {
	IDToken string `json:"id_token"`
}

type GoogleAuthResponse struct {
	AccessToken string `json:"access_token"`
}

func (g GoogleIdToken) ToDomain() *domain.GoogleIdToken {
	return &domain.GoogleIdToken{
		IDToken: g.IDToken,
	}
}

func (g GoogleAuthResponse) FromDomain(dm *domain.GoogleAuthResponse) *GoogleAuthResponse {
	return &GoogleAuthResponse{
		AccessToken: dm.AccessToken,
	}
}
