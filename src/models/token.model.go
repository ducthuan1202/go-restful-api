package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User is struct
type AuthToken struct {
	gorm.Model

	ID           uuid.UUID `gorm:"type:uuid;primaryKey;" json:"id"`
	UserID       uint      `gorm:"type:int;not null" json:"user_id"`
	Token        string    `gorm:"type:text;not null" json:"token"`
	RefreshToken string    `gorm:"type:text;not null" json:"refresh_token"`
	IsActive     bool      `gorm:"type:boolean;not null" json:"is_active"`
}

// https://gorm.io/docs/hooks.html

// TableName
func (t *AuthToken) TableName() string {
	return "auth_tokens"
}

// BeforeCreate is gorm hook
func (t *AuthToken) BeforeCreate(tx *gorm.DB) error {
	t.ID = uuid.New()
	t.IsActive = true
	return nil
}
