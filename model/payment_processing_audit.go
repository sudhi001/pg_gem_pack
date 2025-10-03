package models

import (
	"time"
)

// PaymentProcessingAudit tracks payment processing errors and issues
type PaymentProcessingAudit struct {
	BaseModel
	UID              string     `gorm:"index"`                  // Unique identifier
	PaymentID        string     `gorm:"index"`                  // Payment ID from the message
	TransactionUID   string     `gorm:"index"`                  // Transaction UID from the message
	CustomerID       string     `gorm:"index"`                  // Customer ID from the message
	BranchID         string     `gorm:"index"`                  // Branch ID from the message
	PaymentAmount    float64    `gorm:"index"`                  // Payment amount
	PaymentCurrency  string     `gorm:"index"`                  // Payment currency
	PaymentStatus    string     `gorm:"index"`                  // Payment status from message
	ErrorType        string     `gorm:"index"`                  // Type of error: "CUSTOMER_SYNC", "SCHEME_PROCESSING", "ACME_API", "DATABASE", "VALIDATION", etc.
	ErrorMessage     string     `gorm:"type:text"`              // Detailed error message
	ErrorStack       *string    `gorm:"type:text;default:null"` // Stack trace if available
	ProcessingStage  string     `gorm:"index"`                  // Stage where error occurred: "CUSTOMER_SYNC", "BRANCH_LOOKUP", "SCHEME_CHECK", "INSTALLMENT_ADD", etc.
	RetryCount       int        `gorm:"default:0"`              // Number of retry attempts
	LastRetryAt      *time.Time `gorm:"default:null"`           // Last retry timestamp
	IsResolved       bool       `gorm:"default:false"`          // Whether the issue has been resolved
	ResolutionNote   *string    `gorm:"type:text;default:null"` // Notes about resolution
	ResolvedAt       *time.Time `gorm:"default:null"`           // When the issue was resolved
	RawMessage       string     `gorm:"type:text"`              // Original Kafka message for debugging
	ProcessedAt      time.Time  `gorm:"index"`                  // When the error was recorded
	CustomerPhone    string     `gorm:"index"`                  // Customer phone for reference
	CustomerName     string     `gorm:"index"`                  // Customer name for reference
	GatewayName      string     `gorm:"index"`                  // Payment gateway name
	GatewayOrderID   string     `gorm:"index"`                  // Gateway order ID
	GatewayPaymentID string     `gorm:"index"`                  // Gateway payment ID
}
