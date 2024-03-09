package upload

import (
	"context"
)

type UploadFileProvider struct {
	s3Provider       *s3Provider
	FireBaseProvider *FireBaseProvider
}

func NewUploadFileProvider(s3Provider *s3Provider, fireBaseProvider *FireBaseProvider) *UploadFileProvider {
	return &UploadFileProvider{s3Provider: s3Provider, FireBaseProvider: fireBaseProvider}
}

type UploadProvider interface {
	UploadFile(ctx context.Context, data []byte, dst string) (string, error)
	DeleteFile(objectUrl string) error
	UploadFileFireBase(data []byte, storageObjectPath string) (string, error)
}
