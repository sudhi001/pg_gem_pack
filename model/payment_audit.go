package models

import "time"

// PaymentAuditLog stores all payment-related audit events and conflicts
type PaymentAuditLog struct {
	BaseModel
	TransactionUID string `gorm:"size:100;index;not null"`
	BranchID       string `gorm:"size:50;index;not null"`
	PaymentID      string `gorm:"size:100;index;not null"`
	CustomerID     string `gorm:"size:100;index;not null"`

	// Audit Event Information
	EventType     string `gorm:"size:50;not null"` // CONFLICT, FRAUD_SUSPICION, STATUS_CHANGE, AMOUNT_MISMATCH, etc.
	EventSeverity string `gorm:"size:20;not null"` // LOW, MEDIUM, HIGH, CRITICAL
	EventStatus   string `gorm:"size:20;not null"` // PENDING, INVESTIGATED, RESOLVED, IGNORED

	// Conflict Details
	ConflictType   string `gorm:"size:50"` // STATUS_CONFLICT, AMOUNT_CONFLICT, DUPLICATE_PAYMENT, etc.
	OldValue       JSON   `gorm:"type:jsonb"`
	NewValue       JSON   `gorm:"type:jsonb"`
	ConflictReason string `gorm:"size:500"`

	// Fraud Detection
	FraudScore         float64 `gorm:"default:0"`  // 0-100 fraud probability score
	RiskFactors        JSON    `gorm:"type:jsonb"` // Array of risk factors
	SuspiciousPatterns JSON    `gorm:"type:jsonb"` // Detected suspicious patterns

	// Investigation Details
	InvestigatedBy string `gorm:"size:100"`
	InvestigatedAt time.Time
	Resolution     string `gorm:"size:500"`
	ResolutionTime time.Time

	// Raw Data for Investigation
	RawWebhookData JSON `gorm:"type:jsonb"`
	RawPaymentData JSON `gorm:"type:jsonb"`

	// Email Notification
	EmailSent       bool `gorm:"default:false"`
	EmailSentAt     time.Time
	EmailRecipients string `gorm:"size:500"`

	// Additional Context
	IPAddress string `gorm:"size:45"`
	UserAgent string `gorm:"size:500"`
	DeviceID  string `gorm:"size:100"`
	Location  string `gorm:"size:100"`
}

// PaymentFraudRule defines rules for fraud detection
type PaymentFraudRule struct {
	BaseModel
	RuleName        string `gorm:"size:100;unique;not null"`
	RuleDescription string `gorm:"size:500"`
	RuleType        string `gorm:"size:50;not null"` // AMOUNT_THRESHOLD, FREQUENCY_LIMIT, LOCATION_MISMATCH, etc.
	RuleCategory    string `gorm:"size:50;not null"` // FRAUD, CONFLICT, SUSPICION

	// Rule Configuration
	RuleConfig JSON    `gorm:"type:jsonb;not null"` // Rule-specific configuration
	Threshold  float64 `gorm:"default:0"`
	TimeWindow int     `gorm:"default:0"` // in minutes

	// Rule Status
	IsActive bool `gorm:"default:true"`
	Priority int  `gorm:"default:1"` // 1=Low, 2=Medium, 3=High, 4=Critical

	// Action Configuration
	ActionType    string `gorm:"size:50;not null"` // BLOCK, FLAG, NOTIFY, AUTO_RESOLVE
	EmailTemplate string `gorm:"size:100"`
	AutoResolve   bool   `gorm:"default:false"`

	// Statistics
	TriggerCount    int64 `gorm:"default:0"`
	LastTriggeredAt time.Time
}

// PaymentFraudAlert stores triggered fraud alerts
type PaymentFraudAlert struct {
	BaseModel
	TransactionUID string `gorm:"size:100;index;not null"`
	PaymentID      string `gorm:"size:100;index;not null"`
	BranchID       string `gorm:"size:50;index;not null"`
	CustomerID     string `gorm:"size:100;index;not null"`

	// Alert Information
	RuleID        uint   `gorm:"not null"`
	RuleName      string `gorm:"size:100;not null"`
	AlertType     string `gorm:"size:50;not null"` // FRAUD, SUSPICION, CONFLICT
	AlertSeverity string `gorm:"size:20;not null"` // LOW, MEDIUM, HIGH, CRITICAL
	AlertStatus   string `gorm:"size:20;not null"` // ACTIVE, INVESTIGATED, RESOLVED, FALSE_POSITIVE

	// Fraud Detection Details
	FraudScore     float64 `gorm:"default:0"`
	RiskFactors    JSON    `gorm:"type:jsonb"`
	TriggeredRules JSON    `gorm:"type:jsonb"` // Array of triggered rule IDs

	// Investigation Details
	InvestigatedBy     string `gorm:"size:100"`
	InvestigatedAt     time.Time
	InvestigationNotes string `gorm:"size:1000"`
	Resolution         string `gorm:"size:500"`
	ResolutionTime     time.Time

	// Action Taken
	ActionTaken string `gorm:"size:100"` // BLOCKED, FLAGGED, ALLOWED, REFUNDED
	ActionBy    string `gorm:"size:100"`
	ActionTime  time.Time

	// Notification
	EmailSent      bool `gorm:"default:false"`
	EmailSentAt    time.Time
	SMSAlertSent   bool `gorm:"default:false"`
	SMSAlertSentAt time.Time

	// Context Data
	PaymentAmount    float64 `gorm:"not null"`
	PaymentCurrency  string  `gorm:"size:3;not null"`
	CustomerIP       string  `gorm:"size:45"`
	CustomerLocation string  `gorm:"size:100"`
	DeviceInfo       JSON    `gorm:"type:jsonb"`

	// Raw Data
	RawPaymentData JSON `gorm:"type:jsonb"`
	RawWebhookData JSON `gorm:"type:jsonb"`
}

// PaymentRiskProfile stores customer risk profiles
type PaymentRiskProfile struct {
	BaseModel
	CustomerID string `gorm:"size:100;unique;index;not null"`
	BranchID   string `gorm:"size:50;index;not null"`

	// Risk Assessment
	RiskScore    float64 `gorm:"default:0"`        // 0-100 risk score
	RiskLevel    string  `gorm:"size:20;not null"` // LOW, MEDIUM, HIGH, VERY_HIGH
	RiskCategory string  `gorm:"size:50"`          // NEW_CUSTOMER, HIGH_VALUE, FREQUENT_PAYMENTS, etc.

	// Payment History
	TotalPayments      int64   `gorm:"default:0"`
	SuccessfulPayments int64   `gorm:"default:0"`
	FailedPayments     int64   `gorm:"default:0"`
	TotalAmount        float64 `gorm:"default:0"`
	AverageAmount      float64 `gorm:"default:0"`
	MaxAmount          float64 `gorm:"default:0"`

	// Fraud History
	FraudAlerts     int64 `gorm:"default:0"`
	BlockedPayments int64 `gorm:"default:0"`
	LastFraudAlert  time.Time

	// Location and Device Patterns
	KnownIPs       JSON `gorm:"type:jsonb"` // Array of known IP addresses
	KnownDevices   JSON `gorm:"type:jsonb"` // Array of known device IDs
	KnownLocations JSON `gorm:"type:jsonb"` // Array of known locations

	// Behavioral Patterns
	PaymentPatterns JSON `gorm:"type:jsonb"` // Payment timing patterns
	AmountPatterns  JSON `gorm:"type:jsonb"` // Amount distribution patterns

	// Risk Factors
	RiskFactors JSON `gorm:"type:jsonb"` // Array of risk factors
	RiskHistory JSON `gorm:"type:jsonb"` // Historical risk changes

	// Last Updated
	LastPaymentAt  time.Time
	LastRiskUpdate time.Time
	ProfileStatus  string `gorm:"size:20;default:'ACTIVE'"` // ACTIVE, SUSPENDED, BLOCKED
}

// Constants for Audit Events
const (
	// Event Types
	EventTypeConflict         = "CONFLICT"
	EventTypeFraudSuspicion   = "FRAUD_SUSPICION"
	EventTypeStatusChange     = "STATUS_CHANGE"
	EventTypeAmountMismatch   = "AMOUNT_MISMATCH"
	EventTypeDuplicatePayment = "DUPLICATE_PAYMENT"
	EventTypeLocationMismatch = "LOCATION_MISMATCH"
	EventTypeDeviceMismatch   = "DEVICE_MISMATCH"
	EventTypeHighValuePayment = "HIGH_VALUE_PAYMENT"
	EventTypeFrequencyLimit   = "FREQUENCY_LIMIT"
	EventTypeUnusualPattern   = "UNUSUAL_PATTERN"

	// Event Severity
	EventSeverityLow      = "LOW"
	EventSeverityMedium   = "MEDIUM"
	EventSeverityHigh     = "HIGH"
	EventSeverityCritical = "CRITICAL"

	// Event Status
	EventStatusPending      = "PENDING"
	EventStatusInvestigated = "INVESTIGATED"
	EventStatusResolved     = "RESOLVED"
	EventStatusIgnored      = "IGNORED"

	// Conflict Types
	ConflictTypeStatusConflict   = "STATUS_CONFLICT"
	ConflictTypeAmountConflict   = "AMOUNT_CONFLICT"
	ConflictTypeDuplicatePayment = "DUPLICATE_PAYMENT"
	ConflictTypeCustomerMismatch = "CUSTOMER_MISMATCH"
	ConflictTypeBranchMismatch   = "BRANCH_MISMATCH"

	// Fraud Rule Types
	RuleTypeAmountThreshold  = "AMOUNT_THRESHOLD"
	RuleTypeFrequencyLimit   = "FREQUENCY_LIMIT"
	RuleTypeLocationMismatch = "LOCATION_MISMATCH"
	RuleTypeDeviceMismatch   = "DEVICE_MISMATCH"
	RuleTypeTimePattern      = "TIME_PATTERN"
	RuleTypeAmountPattern    = "AMOUNT_PATTERN"
	RuleTypeCustomerHistory  = "CUSTOMER_HISTORY"
	RuleTypeVelocityCheck    = "VELOCITY_CHECK"

	// Action Types
	ActionTypeBlock        = "BLOCK"
	ActionTypeFlag         = "FLAG"
	ActionTypeNotify       = "NOTIFY"
	ActionTypeAutoResolve  = "AUTO_RESOLVE"
	ActionTypeManualReview = "MANUAL_REVIEW"

	// Risk Levels
	RiskLevelLow      = "LOW"
	RiskLevelMedium   = "MEDIUM"
	RiskLevelHigh     = "HIGH"
	RiskLevelVeryHigh = "VERY_HIGH"
)
