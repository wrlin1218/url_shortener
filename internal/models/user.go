package models

import (
	"crypto/sha512"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID      `gorm:"type:uuid;primary_key;" json:"id"`
	Name      string         `gorm:"unique;not null" json:"name"`
	Password  string         `gorm:"not null" json:"-"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func NewUser(Name, Password string) *User {
	user := &User{
		Name:      Name,
		Password:  Password,
		CreatedAt: time.Now(),
	}
	user.BeforeCreate(nil)
	return user
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	hash := sha512.Sum512([]byte(u.Password))
	u.Password = string(hash[:])
	return nil
}

func (User) TableName() string {
	return "users"
}
