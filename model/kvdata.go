package models

import (
	"time"
)

// KVData represents a key-value data entry
type KVData struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Key       string    `json:"key" gorm:"uniqueIndex;not null"`
	Value     string    `json:"value" gorm:"type:text;not null"`
	Type      string    `json:"type" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName specifies the table name for KVData
func (KVData) TableName() string {
	return "kvdata"
}
