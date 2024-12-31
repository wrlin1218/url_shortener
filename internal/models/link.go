// models/shortlink.go
package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Link struct {
	ID          uuid.UUID      `gorm:"type:uuid;primary_key;" json:"id"`
	UserID      uuid.UUID      `gorm:"type:uuid;index;not null" json:"user_id"`
	OriginalURL string         `gorm:"type:text;not null" json:"original_url"`
	ShortCode   string         `gorm:"unique;not null" json:"short_code"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (s *Link) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.New()
	return
}

func (Link) TableName() string {
	return "links"
}
