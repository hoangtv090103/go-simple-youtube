package port

import (
	"context"

	"gosimpleyoutube/internal/core/domain"
)

type IUserService interface {
	Register(user *domain.User) error
	Login(email, password string) (string, error)
}

type IVideoService interface {
	UploadVideo(video *domain.Video) error
	GetVideo(id string) (*domain.Video, error)
	ListVideos(limit, offset int) ([]*domain.Video, error)
}

type IStorageService interface {
	ProcessVideo(ctx context.Context, videoPath string) (string, error)
}
