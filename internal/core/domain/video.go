package domain

import (
	"time"
)

type Video struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserID      string    `json:"user_id"`
	FileURL     string    `json:"file_url"`
	Thumbnail   string    `json:"thumbnail"`
	Duration    float64   `json:"duration"`
	Views       int64     `json:"views"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
