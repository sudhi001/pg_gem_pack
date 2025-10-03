package models

import (
	"time"
)

// Notification represents a notification sent to a user
type Notification struct {
	BaseModel
	UID          string     `gorm:"index"`
	Title        string     `json:"title" gorm:"type:varchar(255);not null"`
	Body         string     `json:"body" gorm:"type:text;not null"`
	Type         string     `json:"type" gorm:"type:varchar(100);not null"`
	ImageURL     *string    `json:"image_url" gorm:"type:varchar(500)"`
	TotalDevices int        `json:"total_devices" gorm:"default:0"`
	SuccessCount int        `json:"success_count" gorm:"default:0"`
	FailedCount  int        `json:"failed_count" gorm:"default:0"`
	Status       string     `json:"status" gorm:"type:varchar(50);default:'pending'"` // pending, sent, delivered, read
	IsRead       bool       `json:"is_read" gorm:"default:false"`
	DeliveredAt  *time.Time `json:"delivered_at"`
	ReadAt       *time.Time `json:"read_at"`
}

// TableName specifies the table name for Notification
func (Notification) TableName() string {
	return "notifications"
}
