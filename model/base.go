package models

import "time"

// BaseModel holds common fields for all models
type BaseModel struct {
	ID           uint      `gorm:"primaryKey;autoIncrement;index"`
	SpaceUID     string    `gorm:"size:50;null"`
	CreatedByUID string    `gorm:"index;null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}

// JSON type for handling JSON data in Go
type JSON map[string]interface{}
