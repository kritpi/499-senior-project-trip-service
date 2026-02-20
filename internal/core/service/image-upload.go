package service

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
)

type supabaseSignedURLResponse struct {
	SignedURL string `json:"signedURL"`
}

func normalizeSignedURL(baseURL, signedPath string) string {
	// Supabase sometimes returns /object/sign/...
	if strings.HasPrefix(signedPath, "/object/") {
		return baseURL + "/storage/v1" + signedPath
	}

	// Normal case
	if strings.HasPrefix(signedPath, "/storage/") {
		return baseURL + signedPath
	}

	return baseURL + signedPath
}

func (s *service) UploadImage(
	ctx context.Context,
	in domain.UploadImageRequest,
) (*domain.UploadImageResponse, error) {

	file, err := in.File.Open()
	if err != nil {
		return nil, fmt.Errorf("open file: %w", err)
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}

	timestamp := time.Now().Unix()
	ext := filepath.Ext(in.File.Filename)
	filename := fmt.Sprintf("%s_%d%s", in.MemberId, timestamp, ext)
	filePath := fmt.Sprintf("images/%s", filename)

	uploadEndpoint := fmt.Sprintf(
		"%s/storage/v1/object/%s/%s",
		s.cfg.Storage.Url,
		s.cfg.Storage.Bucket,
		filePath,
	)

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		uploadEndpoint,
		bytes.NewReader(fileBytes),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+s.cfg.Storage.ApiKey)
	req.Header.Set("Content-Type", in.ContentType)
	req.Header.Set("x-upsert", "true")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("upload failed (%d): %s", resp.StatusCode, body)
	}

	// ✅ PUBLIC URL (NO SIGNING)
	publicURL := fmt.Sprintf(
		"%s/storage/v1/object/public/%s/%s",
		s.cfg.Storage.Url,
		s.cfg.Storage.Bucket,
		filePath,
	)

	return &domain.UploadImageResponse{
		FilePath: filePath,
		Url:      publicURL,
	}, nil
}
