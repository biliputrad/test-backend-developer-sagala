package base_entities

import (
	"gorm.io/gorm"
	"time"
)

type Base struct {
	ID        int64          `gorm:"type:uuid;primaryKey" json:"id"`
	CreatedAt time.Time      `gorm:"<-create" json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	CreatedBy int64          `json:"created_by"`
	UpdatedBy int64          `json:"updated_by"`
	DeletedBy int64          `json:"deleted_by"`
}
