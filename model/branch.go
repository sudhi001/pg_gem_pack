package models

import "github.com/lib/pq"

// Branche model for handling payment details

type Branch struct {
	BaseModel
	UID                        string         `gorm:"index"`              // Unique identifier, indexed
	Name                       string         `gorm:"index"`              // Name of the branch, indexed
	Code                       string         `gorm:"index"`              // Branch code, not indexed
	ConfigID                   string         `gorm:"index"`              // Configuration ID for the branch
	Description                string         `gorm:"index"`              // Description of the branch, not indexed
	GST                        string         `gorm:"index"`              // GST details, not indexed
	Address                    string         `gorm:"index"`              // Branch address, not indexed
	Status                     *string        `gorm:"default:'RELEASED'"` // Branch status: TESTING, RELEASED, SUSPENDED
	AppPaymentGatewayName      *string        `gorm:"default:null"`       // Name of the payment gateway
	AppPaymentGatewayClient    *string        `gorm:"default:null"`       // Payment gateway client details
	AppPaymentGatewaySecret    *string        `gorm:"default:null"`       // Payment gateway secret
	AppPaymentGatewayEnv       *string        `gorm:"default:null"`       // Payment gateway secret
	PaymentType                *string        `gorm:"default:null"`       // Payment gateway type: "hdfc_smart_gateway" or "razorpay"
	RazorpayKeyId              *string        `gorm:"default:null"`       // Razorpay Key ID
	RazorpayKeySecret          *string        `gorm:"default:null"`       // Razorpay Key Secret (encrypted)
	RazorpayKeyIdForBeta       *string        `gorm:"default:null"`       // Beta/testing Razorpay Key ID
	RazorpayKeySecretForBeta   *string        `gorm:"default:null"`       // Beta/testing Razorpay Key Secret (encrypted)
	TestBuildNumber            *int           `gorm:"default:null"`       // Build number threshold for using beta credentials
	PaymentAccepted            *bool          `gorm:"default:true"`       // Whether payment is accepted for this branch
	PaymentGateWayErrorMessage *string        `gorm:"default:null"`       // Error message when payment is not accepted
	Office                     bool           `gorm:"default:false"`      // Flag to identify if this is an office branch
	PhoneNumbers               pq.StringArray `gorm:"type:text[]"`        // Array of phone numbers (PostgreSQL ARRAY type)
}
