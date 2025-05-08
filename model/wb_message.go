package models

import (
	"encoding/json"
	"time"

	"github.com/lib/pq"
)

// WhatsApp represents the 'whatsapp' table
type WhatsApp struct {
	BaseModel
	Timestamp time.Time       `gorm:"not null" json:"timestamp"`
	Data      json.RawMessage `json:"data" gorm:"type:json"`
}

// WhatsAppBusinessMessageRequest represents the 'wb_message_requests' table
type WhatsAppBusinessMessageRequest struct {
	BaseModel
	UID              string                                   `gorm:"not null;index"`
	TemplateName     string                                   `gorm:"size:100;" json:"template_name"`
	Data             json.RawMessage                          `gorm:"not null"`
	PhoneNumbers     pq.StringArray                           `gorm:"type:text[]"`
	Responses        pq.StringArray                           `gorm:"type:jsonb[]"` // Change this line
	ReferenceID      string                                   `gorm:"index"`
	MessageResponses []WhatsAppBusinessMessageRequestResponse `gorm:"foreignKey:RequestID;references:ID;constraint:OnDelete:CASCADE"`
	SentCount        int                                      `gorm:"default:0" json:"sent_count"`
	DeliveredCount   int                                      `gorm:"default:0" json:"delivered_count"`
	ReadCount        int                                      `gorm:"default:0" json:"read_count"`
	FailedCount      int                                      `gorm:"default:0" json:"failed_count"`
	ExpiredCount     int                                      `gorm:"default:0" json:"expired_count"`
	RepliedCount     int                                      `gorm:"default:0" json:"replied_count"`
	BlockedCount     int                                      `gorm:"default:0" json:"blocked_count"`
}

// WhatsAppBusinessMessageRequestResponse represents the 'wb_message_requests_response' table
type WhatsAppBusinessMessageRequestResponse struct {
	BaseModel
	RecipientID string    `gorm:"not null"`
	Timestamp   time.Time `gorm:"not null"`
	Status      string    `gorm:"not null"`
	PhoneNumber string    `gorm:"not null"`
	RequestID   uint      `gorm:"not null;index;"`
}

// TableName sets custom table names for the models
func (WhatsApp) TableName() string {
	return "whatsapp"
}
func (WhatsAppBusinessMessageRequestResponse) TableName() string {
	return "wb_message_requests_response"
}
func (WhatsAppBusinessMessageRequest) TableName() string {
	return "wb_message_requests"
}
