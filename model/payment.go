package models

import "time"

// Payment model for storing payment information
type Payment struct {
	BaseModel
	TransactionUID string `gorm:"size:100;index;not null"` // TransactionUID from webhook
	BranchID       string `gorm:"size:50;index;not null"`  // Branch ID from webhook

	// Product and Branch Information
	ProductCode string `gorm:"size:50"`  // Product/Scheme code
	ProductName string `gorm:"size:255"` // Product/Scheme code
	BranchCode  string `gorm:"size:50"`  // Branch code
	BranchUID   string `gorm:"size:255"` // Branch UID
	UID         string `gorm:"size:255"` // Subscribable UID

	// Order Information
	OrderID       string  `gorm:"size:100;index;not null"`
	OrderAmount   float64 `gorm:"not null"`
	OrderCurrency string  `gorm:"size:3;not null"`
	OrderTags     JSON    `gorm:"type:jsonb"`

	// Payment Information
	PaymentID       string  `gorm:"size:100;uniqueIndex;not null"` // cf_payment_id
	PaymentStatus   string  `gorm:"size:50;not null"`              // SUCCESS, FAILED, PENDING
	PaymentAmount   float64 `gorm:"not null"`
	PaymentCurrency string  `gorm:"size:3;not null"`
	PaymentMessage  string  `gorm:"size:500"`
	PaymentTime     time.Time
	BankReference   string `gorm:"size:100"`
	AuthID          string `gorm:"size:100"`

	// Payment Method Details
	PaymentMethod JSON   `gorm:"type:jsonb"` // UPI, Card, etc. details
	PaymentGroup  string `gorm:"size:50"`    // upi, card, netbanking, etc.

	// Customer Information
	CustomerID    string `gorm:"size:100;index;not null"`
	CustomerName  string `gorm:"size:200"`
	CustomerEmail string `gorm:"size:200"`
	CustomerPhone string `gorm:"size:20"`

	// Payment Gateway Details
	GatewayName             string `gorm:"size:50"`
	GatewayOrderID          string `gorm:"size:100"`
	GatewayPaymentID        string `gorm:"size:100"`
	GatewayStatusCode       string `gorm:"size:50"`
	GatewayOrderReferenceID string `gorm:"size:100"`
	GatewaySettlement       string `gorm:"size:50"`
	GatewayReferenceName    string `gorm:"size:100"`

	// Webhook Information
	WebhookType      string `gorm:"size:50;not null"` // PAYMENT_SUCCESS_WEBHOOK, PAYMENT_FAILED_WEBHOOK, etc.
	WebhookEventTime time.Time
	WebhookSignature string `gorm:"size:500"`
	WebhookTimestamp string `gorm:"size:50"`
	WebhookVersion   string `gorm:"size:20"`

	// Idempotency and Attempts
	IdempotencyKey string `gorm:"size:500"`
	Attempt        int    `gorm:"default:1"`

	// Notification tracking
	NotificationSent bool `gorm:"default:false"` // Track if FCM notification has been sent

	// Raw webhook data for debugging
	RawWebhookData JSON `gorm:"type:jsonb"`
	Payload        JSON `gorm:"type:jsonb;not null"`

	// Relationship to Payment Charges
	PaymentCharges []PaymentCharge `gorm:"foreignKey:PaymentID;references:PaymentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// TableName specifies the table name for Payment
func (Payment) TableName() string {
	return "payments"
}

// TableName specifies the table name for PaymentCharge
func (PaymentCharge) TableName() string {
	return "payment_charges"
}

// PaymentCharge model for storing payment charges information
type PaymentCharge struct {
	BaseModel
	PaymentID string `gorm:"size:100;index;not null"` // Foreign key to Payment.PaymentID

	// Charges Details
	ServiceCharge         float64 `gorm:"not null"`
	ServiceTax            float64 `gorm:"not null"`
	SettlementAmount      float64 `gorm:"not null"`
	SettlementCurrency    string  `gorm:"size:3;not null"`
	ServiceChargeDiscount float64

	// Webhook Information
	WebhookType      string `gorm:"size:50;not null"` // PAYMENT_CHARGES_WEBHOOK
	WebhookEventTime time.Time
	WebhookSignature string `gorm:"size:500"`
	WebhookTimestamp string `gorm:"size:50"`
	WebhookVersion   string `gorm:"size:20"`

	// Raw webhook data for debugging
	RawWebhookData JSON `gorm:"type:jsonb"`
}

// PaymentStatus constants
const (
	PaymentStatusSuccess  = "SUCCESS"
	PaymentStatusFailed   = "FAILED"
	PaymentStatusPending  = "PENDING"
	PaymentStatusRefunded = "REFUNDED"
)

// WebhookType constants
const (
	WebhookTypePaymentSuccess = "PAYMENT_SUCCESS_WEBHOOK"
	WebhookTypePaymentFailed  = "PAYMENT_FAILED_WEBHOOK"
	WebhookTypePaymentCharges = "PAYMENT_CHARGES_WEBHOOK"
	WebhookTypeRefundSuccess  = "REFUND_SUCCESS_WEBHOOK"
	WebhookTypeRefundFailed   = "REFUND_FAILED_WEBHOOK"
)

// PaymentGroup constants
const (
	PaymentGroupUPI        = "upi"
	PaymentGroupCard       = "card"
	PaymentGroupNetBanking = "netbanking"
	PaymentGroupWallet     = "wallet"
	PaymentGroupEMI        = "emi"
)

// GatewayName constants
const (
	GatewayNameCashfree = "CASHFREE"
	GatewayNameRazorpay = "RAZORPAY"
	GatewayNamePayU     = "PAYU"
	GatewayNameStripe   = "STRIPE"
)
