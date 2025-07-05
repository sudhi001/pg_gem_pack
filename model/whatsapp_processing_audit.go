package models

import (
	"time"
)

// WhatsAppProcessingAudit tracks WhatsApp message processing errors and issues
type WhatsAppProcessingAudit struct {
	BaseModel
	UID              string     `gorm:"index"`                  // Unique identifier
	MessageID        string     `gorm:"index"`                  // WhatsApp message ID
	PhoneNumber      string     `gorm:"index"`                  // Phone number from the message
	CustomerID       string     `gorm:"index"`                  // Customer ID if available
	MessageType      string     `gorm:"index"`                  // Type of message: "text", "media", "template", "webhook"
	ErrorType        string     `gorm:"index"`                  // Type of error: "PARSING", "VALIDATION", "WEBHOOK", "MEDIA", "TEMPLATE", "DATABASE", etc.
	ErrorMessage     string     `gorm:"type:text"`              // Detailed error message
	ErrorStack       *string    `gorm:"type:text;default:null"` // Stack trace if available
	ProcessingStage  string     `gorm:"index"`                  // Stage where error occurred: "MESSAGE_PARSING", "WEBHOOK_PROCESSING", "MEDIA_DOWNLOAD", "TEMPLATE_SENDING", etc.
	RetryCount       int        `gorm:"default:0"`              // Number of retry attempts
	LastRetryAt      *time.Time `gorm:"default:null"`           // Last retry timestamp
	IsResolved       bool       `gorm:"default:false"`          // Whether the issue has been resolved
	ResolutionNote   *string    `gorm:"type:text;default:null"` // Notes about resolution
	ResolvedAt       *time.Time `gorm:"default:null"`           // When the issue was resolved
	RawMessage       string     `gorm:"type:text"`              // Original Kafka message for debugging
	ProcessedAt      time.Time  `gorm:"index"`                  // When the error was recorded
	TopicName        string     `gorm:"index"`                  // Kafka topic name
	TemplateName     string     `gorm:"index"`                  // Template name if applicable
	MediaType        string     `gorm:"index"`                  // Media type if applicable
	WebhookType      string     `gorm:"index"`                  // Webhook type if applicable
	Status           string     `gorm:"index"`                  // Message status if available
}

// TableName specifies the table name for WhatsAppProcessingAudit
func (WhatsAppProcessingAudit) TableName() string {
	return "whatsapp_processing_audit"
} 