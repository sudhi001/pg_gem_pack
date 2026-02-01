package models

import "time"

// AppLog stores security and unusual events for developer monitoring
type AppLog struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	SpaceUID    string    `json:"space_uid" gorm:"size:50;index"`
	Level       string    `json:"level" gorm:"size:20;index"`                      // INFO, WARNING, ERROR, SECURITY
	EventType   string    `json:"event_type" gorm:"size:50;index"`                 // DEVICE_MISMATCH, DEVICE_OWNERSHIP_DENIED, etc.
	Message     string    `json:"message" gorm:"type:text"`                        // Human-readable message for developers
	UserID      string    `json:"user_id" gorm:"size:50;index"`                    // Affected user ID
	DeviceID    string    `json:"device_id" gorm:"size:50;index"`                  // Related device ID
	RequestID   string    `json:"request_id" gorm:"size:50;index"`                 // Request ID for tracing
	IPAddress   string    `json:"ip_address" gorm:"size:50"`                       // Client IP
	UserAgent   string    `json:"user_agent" gorm:"type:text"`                     // User agent string
	Endpoint    string    `json:"endpoint" gorm:"size:100"`                        // API endpoint
	ExtraData   JSON      `json:"extra_data" gorm:"type:jsonb"`                    // Additional context as JSON
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime;index"`
}

// TableName specifies the table name for AppLog
func (AppLog) TableName() string {
	return "app_logs"
}

// Log levels
const (
	LogLevelInfo     = "INFO"
	LogLevelWarning  = "WARNING"
	LogLevelError    = "ERROR"
	LogLevelSecurity = "SECURITY"
)

// Event types
const (
	EventDeviceMismatch        = "DEVICE_MISMATCH"
	EventDeviceOwnershipDenied = "DEVICE_OWNERSHIP_DENIED"
	EventInvalidOTP            = "INVALID_OTP"
	EventOTPExpired            = "OTP_EXPIRED"
	EventVerificationNotFound  = "VERIFICATION_NOT_FOUND"
	EventUnauthorizedAccess    = "UNAUTHORIZED_ACCESS"
	EventSuspiciousActivity    = "SUSPICIOUS_ACTIVITY"
)
