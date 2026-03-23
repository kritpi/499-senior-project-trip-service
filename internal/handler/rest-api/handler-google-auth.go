package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
	"github.com/kritpi/499-senior-project-trip-service/internal/handler/dto"
)

func (h *restHandler) GoogleAuth(c *fiber.Ctx) error {
	ctx := c.Context()
	var req dto.GoogleAuthRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	// Exchange code for Google ID token
	idToken, err := OutboundExchangeOAuthToken(
		req.Code,
		req.CodeVerifier,
		req.RedirectURI,
		h.cfg.Auth.ClientId,
		h.cfg.Auth.ClientSecrets,
	)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	tokenResp, err := h.svc.GoogleAuth(ctx, domain.GoogleIdToken{
		IDToken: idToken,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.GoogleAuthResponse{
		AccessToken: tokenResp.AccessToken,
	})
}

// OutboundExchangeOAuthToken exchanges an authorization code for a Google ID token
func OutboundExchangeOAuthToken(code, codeVerifier, redirectURI, clientID, clientSecret string) (string, error) {
	tokenEndpoint := "https://oauth2.googleapis.com/token"
	data := url.Values{}
	data.Set("code", code)
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)
	data.Set("redirect_uri", redirectURI)
	data.Set("grant_type", "authorization_code")
	data.Set("code_verifier", codeVerifier)

	resp, err := http.PostForm(tokenEndpoint, data)
	if err != nil {
		return "", fmt.Errorf("failed to exchange token: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		log.Printf("Google Token Exchange Failed. Status: %d, Body: %s", resp.StatusCode, string(bodyBytes))

		var errBody map[string]interface{}
		_ = json.Unmarshal(bodyBytes, &errBody)
		return "", fmt.Errorf("google token exchange failed: %v", errBody)
	}

	var googleToken struct {
		IdToken string `json:"id_token"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&googleToken); err != nil {
		return "", fmt.Errorf("failed to parse google response: %w", err)
	}

	return googleToken.IdToken, nil
}
