package dto

import (
	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
	"github.com/kritpi/499-senior-project-trip-service/shared/model"
)

type JWTCustomClaims model.JWTCustomClaims

type GoogleAuthRequest struct {
	Code         string `json:"code"`
	CodeVerifier string `json:"code_verifier"`
	RedirectURI  string `json:"redirect_uri"`
}

type GoogleAuthResponse struct {
	AccessToken string `json:"access_token"`
}

func (g GoogleAuthResponse) FromDomain(dm *domain.GoogleAuthResponse) *GoogleAuthResponse {
	return &GoogleAuthResponse{
		AccessToken: dm.AccessToken,
	}
}
