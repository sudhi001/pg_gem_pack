package models

import (
	"time"
)

// CustomerAudit tracks customer creation and synchronization with external providers
type CustomerAudit struct {
	BaseModel
	UID              string    `gorm:"index"`                  // Unique identifier
	UserID           string    `gorm:"index"`                  // Our user ID
	CustomerPhone    string    `gorm:"index"`                  // Customer phone number
	CustomerName     string    `gorm:"index"`                  // Customer name
	ProviderCode     int       `gorm:"index"`                  // External provider customer code
	ProviderName     string    `gorm:"index"`                  // Provider name (e.g., "ACME", "OTHER_PROVIDER")
	Action           string    `gorm:"index"`                  // Action performed: "CREATED", "FOUND", "UPDATED"
	ProviderResponse string    `gorm:"type:text"`              // Full provider API response
	ErrorMessage     *string   `gorm:"type:text;default:null"` // Error message if any
	ProcessedAt      time.Time `gorm:"index"`                  // When the action was processed
	PaymentID        *string   `gorm:"index;default:null"`     // Associated payment ID if any
	TransactionUID   *string   `gorm:"index;default:null"`     // Associated transaction UID if any
}
