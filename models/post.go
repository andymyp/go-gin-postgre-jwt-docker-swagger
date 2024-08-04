package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	ID        uuid.UUID `gorm:"primaryKey" json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	User      User      `gorm:"references:ID" json:"user"`
	Title     string    `gorm:"type:varchar(150)" json:"title" validate:"required"`
	Content   string    `gorm:"type:text" json:"content" validate:"required"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (post *Post) BeforeCreate(tx *gorm.DB) (err error) {
	post.ID = uuid.New()
	return nil
}

type PostResponse struct {
	ID        uuid.UUID    `json:"id"`
	User      UserResponse `json:"user"`
	Title     string       `json:"title"`
	Content   string       `json:"content"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}

type InputPost struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}
