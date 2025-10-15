package models

import (
	"time"
)

// FAQ represents a frequently asked question entry
type FAQ struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	Question   string    `json:"question" gorm:"type:text;not null"`
	Answer     string    `json:"answer" gorm:"type:text;not null"`
	Category   string    `json:"category" gorm:"type:varchar(100);default:'General'"`
	IsActive   bool      `json:"is_active" gorm:"default:true"`
	SortOrder  int       `json:"sort_order" gorm:"default:0"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// TableName returns the table name for the FAQ model
func (FAQ) TableName() string {
	return "faqdata"
}

// FAQResponse represents the API response structure for FAQ data
type FAQResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    []FAQ  `json:"data"`
	Count   int    `json:"count"`
}

// FAQItemResponse represents a single FAQ item response
type FAQItemResponse struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Data    FAQ  `json:"data"`
}
