package postgres

import (
	"database/sql"

	"gosimpleyoutube/internal/core/domain"

	"github.com/jmoiron/sqlx"
)

type postgresVideoRepository struct {
	db *sql.DB
}

func NewPostgresVideoRepository(db *sql.DB) *postgresVideoRepository {
	return &postgresVideoRepository{
		db: db,
	}
}

func (r *postgresVideoRepository) Store(video *domain.Video) error {
	query := `
		INSERT INTO videos (id, title, description, user_id, file_url, thumbnail, duration, views, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	query = sqlx.Rebind(sqlx.DOLLAR, query)

	_, err := r.db.Exec(
		query,
		video.ID,
		video.Title,
		video.Description,
		video.UserID,
		video.FileURL,
		video.Thumbnail,
		video.Duration,
		video.Views,
		video.CreatedAt,
		video.UpdatedAt,
	)

	return err
}

func (r *postgresVideoRepository) FindByID(id string) (*domain.Video, error) {
	return nil, nil
}

func (r *postgresVideoRepository) FindAll(limit, offset int) ([]*domain.Video, error) {
	return nil, nil
}
