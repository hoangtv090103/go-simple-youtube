package service

import (
	"io"
	"time"
)

type StorageService interface {
	UploadFile(fileName string, fileContent io.Reader) (string, error)
	DeleteFile(fileURL string) error
	GetSignedURL(fileURL string, expiration time.Duration) (string, error)
}
