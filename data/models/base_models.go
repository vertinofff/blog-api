package models

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	Id        int       `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"type:TIMESTAMPTZ;not null"`
	Version   uint      `gorm:"not null;default:1"`
}

func (m *BaseModel) BeforeCreate(tx *gorm.DB) error { m.CreatedAt = time.Now().UTC(); return nil }
