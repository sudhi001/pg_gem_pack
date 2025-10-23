package models

import "time"

// WhatsAppContact represents WhatsApp contact information in the database
type WhatsAppContact struct {
	BaseModel
	// Contact identifiers
	WaId string `gorm:"uniqueIndex:idx_waid;type:varchar(20);not null" json:"wa_id"`
	Name string `gorm:"type:varchar(100)" json:"name"`

	// Metadata
	PhoneNumber    string `gorm:"index:idx_phone;type:varchar(20)" json:"display_phone_number"`
	PhoneNumberId  string `gorm:"type:varchar(50);not null" json:"phone_number_id"`
	ProfilePicture string `gorm:"type:varchar(500)" json:"profile_picture"`

	// Tracking fields
	LastSeen time.Time `gorm:"index" json:"last_seen"`
	Active   bool      `gorm:"default:true" json:"active"`
}

// Add this method to your WhatsAppContact struct
func (WhatsAppContact) TableName() string {
	return "whats_app_contacts" // Note the underscores
}
