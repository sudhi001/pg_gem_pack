package models

import (
	"time"

	"gorm.io/gorm"
)

// SecureTransaction model for handling payment details
type SecureTransaction struct {
	BaseModel
	SecureUID  string    `gorm:"index"`            // Unique identifier, indexed
	Name       string    `gorm:"index"`            // Name, indexed
	PublicKEY  string    `gorm:"default:not null"` // Public key
	PrivateKEY string    `gorm:"default:not null"` // Private key
	ExpireON   time.Time `gorm:"autoCreateTime"`   // Expiration time (in UTC)
}

// BeforeCreate is a GORM hook that runs before creating a SecureTransaction
func (s *SecureTransaction) BeforeCreate(tx *gorm.DB) (err error) {
	// Set the ExpireOn field to the current UTC time + 365 days
	s.ExpireON = time.Now().UTC().AddDate(1, 0, 0)
	return nil
}
