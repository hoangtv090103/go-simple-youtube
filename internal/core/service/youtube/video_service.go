package service

import (
	"gosimpleyoutube/internal/core/domain"
	"gosimpleyoutube/internal/core/port"
)

type VideoService struct {
	videoRepo port.IVideoRepository
}

func NewVideoService(videoRepo port.IVideoRepository) *VideoService {
	return &VideoService{
		videoRepo: videoRepo,
	}
}

func (uc *VideoService) UploadVideo(video *domain.Video) error {
	return uc.videoRepo.Store(video)
}

func (uc *VideoService) GetVideo(id string) (*domain.Video, error) {
	return uc.videoRepo.FindByID(id)
}

func (uc *VideoService) ListVideos(limit, offset int) ([]*domain.Video, error) {
	return uc.videoRepo.FindAll(limit, offset)
}
