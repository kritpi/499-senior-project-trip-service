package domain

import "mime/multipart"

type UploadImageRequest struct {
	File        *multipart.FileHeader
	ContentType string
	MemberId    string
}

type UploadImageResponse struct {
	Url      string
	FilePath string
}
