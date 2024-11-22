package port

import "gosimpleyoutube/internal/core/domain"

type IUserRepository interface {
	Store(user *domain.User) error
	FindByID(id string) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	Update(user *domain.User) error
	Delete(id string) error
}

type IVideoRepository interface {
	Store(video *domain.Video) error
	FindByID(id string) (*domain.Video, error)
	FindAll(limit, offset int) ([]*domain.Video, error)
}
