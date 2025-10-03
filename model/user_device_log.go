package models

import (
	"time"
)

// UserDeviceLog represents user device login/logout sessions
type UserDeviceLog struct {
	BaseModel
	SessionID        string     `gorm:"index;not null"`                    // Session identifier
	UserID           string     `gorm:"index;not null"`                    // User ID who logged in
	PhoneNumber      string     `gorm:"index;not null"`                    // Phone number used for login
	RealIP           string     `gorm:"index;not null"`                    // Real IP address of the user
	UserAgent        string     `gorm:"type:text;not null"`                // Browser user agent string
	BrowserName      string     `gorm:"index"`                             // Browser name (Chrome, Firefox, etc.)
	BrowserVersion   string     `gorm:"index"`                             // Browser version
	OSName           string     `gorm:"index"`                             // Operating system name
	OSVersion        string     `gorm:"index"`                             // Operating system version
	DeviceType       string     `gorm:"index"`                             // Device type (desktop, mobile, tablet)
	LoginTime        time.Time  `gorm:"index;not null"`                    // Time when user logged in
	LogoutTime       *time.Time `gorm:"index;default:null"`                // Time when user logged out (null if still active)
	IsActive         bool       `gorm:"default:true;index"`                // Whether the session is still active
	LoginLocation    *string    `gorm:"type:geometry(Point,4326);default:null"` // Geographic location at login (SRID 4326)
	LogoutLocation   *string    `gorm:"type:geometry(Point,4326);default:null"` // Geographic location at logout
	LoginReason      string     `gorm:"default:'manual'"`                  // Reason for login (manual, auto, etc.)
	LogoutReason     string     `gorm:"default:'manual'"`                  // Reason for logout (manual, timeout, etc.)
	SessionDuration  *int64     `gorm:"default:null"`                      // Session duration in seconds (calculated when logged out)
	AdditionalData   JSON       `gorm:"type:jsonb;default:null"`           // Additional session data as JSON
} 