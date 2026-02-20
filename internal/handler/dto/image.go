package dto

import "github.com/kritpi/499-senior-project-trip-service/internal/core/domain"

type UploadImageResponse struct {
	Url      string `json:"image_url"`
	FilePath string `json:"file_path"`
}

func (UploadImageResponse) FromDomain(d *domain.UploadImageResponse) UploadImageResponse {
	return UploadImageResponse{
		Url:      d.Url,
		FilePath: d.FilePath,
	}
}
