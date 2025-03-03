package models

import "time"

// Branche model for handling payment details

type PhoneNumberVerifications struct {
	BaseModel
	VerificationID       string     `gorm:"index"`         // Verification ID
	UserID               string     `gorm:"index"`         // User ID
	DeviceID             string     `gorm:"index"`         // Device ID
	PhoneNumber          string     `gorm:"index"`         // Phone number
	PhoneNumberExtension string     `gorm:"index"`         // Phone number extension
	OTP                  string     `gorm:"index"`         // OTP
	Verified             *bool      `gorm:"default:null"`  // Verified status (nullable)
	CMD                  *string    `gorm:"default:null"`  // Command (e.g., 0 - VERIFICATION, 1 - ACCOUNT DELETION)
	OTPGeneratedOn       *time.Time `gorm:"default:null"`  // Time when the OTP was generated
	Send                 bool       `gorm:"default:false"` // Indicates if OTP was sent
	StatusMessage        *string    `gorm:"default:null"`  // Status message
	RequestHash          string     `gorm:"index"`
}
