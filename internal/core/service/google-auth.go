package service

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"

	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
	"github.com/kritpi/499-senior-project-trip-service/shared/utils"
	"google.golang.org/api/idtoken"
)

func (s *service) GoogleAuth(ctx context.Context, idToken domain.GoogleIdToken) (*domain.GoogleAuthResponse, error) {
	now := time.Now().Local()
	memberId := uuid.New()

	googleClientId := s.cfg.Auth.ClientId
	accountPayload, err := VerifyGoogleIdToken(ctx, idToken.IDToken, googleClientId)
	if err != nil || accountPayload == nil {
		log.Errorf("unable to validate google id token: %+v", err)
		return nil, err
	}

	// query member with email to check if member exists
	member, err := s.repo.GetMemberByEmail(ctx, accountPayload.Email)
	if err != nil {
		log.Errorf("unable to get member form db: %+v", err)
		return nil, err
	}
	if member == nil {
		// create new member
		err := s.repo.CreateMember(ctx, domain.Member{
			ID:       memberId.String(),
			Name:     accountPayload.Name,
			Email:    accountPayload.Email,
			ImageUrl: accountPayload.ImageUrl,
		})
		if err != nil {
			log.Errorf("unable to create new member: +%v", err)
			return nil, err
		}
	}

	// Issue JWT
	token, err := utils.GenerateToken(*accountPayload, now, s.cfg)
	if err != nil {
		return nil, err
	}

	resp := domain.GoogleAuthResponse{
		AccessToken: *token,
	}

	return &resp, nil
}

func VerifyGoogleIdToken(ctx context.Context, googleIdToken, googleClientID string) (*domain.GoogleUserClaims, error) {
	payload, err := idtoken.Validate(ctx, googleIdToken, googleClientID)
	if err != nil {
		return nil, err
	}

	getString := func(key string) string {
		if v, ok := payload.Claims[key].(string); ok {
			return v
		}
		return ""
	}

	user := domain.GoogleUserClaims{
		Sub:      getString("sub"),
		Email:    getString("email"),
		Name:     getString("name"),
		ImageUrl: getString("picture"),
	}

	return &user, nil
}
