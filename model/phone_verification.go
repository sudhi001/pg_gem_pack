package models

import (
	"time"
)

// PhoneVerification represents phone verification data in the system
type PhoneVerification struct {
	BaseModel
	VerificationID       string     `gorm:"column:verification_id"`
	UserID               string     `gorm:"column:user_id"`
	DeviceID             string     `gorm:"column:device_id"`
	PhoneNumber          string     `gorm:"column:phone_number"`
	PhoneNumberExtension string     `gorm:"column:phone_number_extension"`
	OTP                  string     `gorm:"column:otp"`
	Verified             *bool      `gorm:"column:verified;default:null"`
	CMD                  *string    `gorm:"column:cmd;default:null"` // 0 - VERIFICATION, 1 - ACCOUNT DELETION
	OTPGeneratedOn       *time.Time `gorm:"column:otp_generated_on;default:null"`
	Send                 bool       `gorm:"column:send;default:false"`
	StatusMessage        *string    `gorm:"column:status_message;default:null"`
	RequestHash          string     `gorm:"column:requestHash;index"`
}

// IsOTPExpired checks if the OTP has expired based on the given expiration duration
func (p *PhoneVerification) IsOTPExpired(expirationMinutes int) bool {
	if p.OTPGeneratedOn == nil {
		return true // OTP not generated
	}

	currentTime := time.Now().UTC()
	expirationThreshold := time.Duration(expirationMinutes) * time.Minute

	// Check if the time difference is greater than the expiration threshold
	return currentTime.Sub(*p.OTPGeneratedOn) > expirationThreshold
}
