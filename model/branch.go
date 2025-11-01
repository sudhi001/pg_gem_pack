package models

// Branche model for handling payment details

type Branch struct {
	BaseModel
	UID                     string   `gorm:"index"`              // Unique identifier, indexed
	Name                    string   `gorm:"index"`              // Name of the branch, indexed
	Code                    string   `gorm:"index"`              // Branch code, not indexed
	Description             string   `gorm:"index"`              // Description of the branch, not indexed
	GST                     string   `gorm:"index"`              // GST details, not indexed
	Address                 string   `gorm:"index"`              // Branch address, not indexed
	Status                  *string  `gorm:"default:'RELEASED'"` // Branch status: TESTING, RELEASED, SUSPENDED
	AppPaymentGatewayName   *string  `gorm:"default:null"`       // Name of the payment gateway
	AppPaymentGatewayClient *string  `gorm:"default:null"`       // Payment gateway client details
	AppPaymentGatewaySecret *string  `gorm:"default:null"`       // Payment gateway secret
	AppPaymentGatewayEnv    *string  `gorm:"default:null"`       // Payment gateway secret
	PaymentType             *string  `gorm:"default:null"`       // Payment gateway type: "hdfc_smart_gateway" or "razorpay"
	RazorpayKeyId           *string  `gorm:"default:null"`       // Razorpay Key ID
	RazorpayKeySecret       *string  `gorm:"default:null"`       // Razorpay Key Secret (encrypted)
	RazorpayWebhookSecret   *string  `gorm:"default:null"`       // Razorpay Webhook Secret (encrypted)
	Office                  bool     `gorm:"default:false"`      // Flag to identify if this is an office branch
	PhoneNumbers            []string `gorm:"type:text[]"`        // Array of phone numbers (PostgreSQL ARRAY type)
}
