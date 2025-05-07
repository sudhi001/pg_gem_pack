package models

import "time"

// Create WWhatsAppMessageStatus struct
// WhatsAppMessageStatus represents the status of a WhatsApp message
type WhatsAppMessageStatus struct {
	BaseModel
	MessageID               string `gorm:"index"`
	RecipientID             string `gorm:"index"`
	Status                  string 
	Timestamp               time.Time
	PhoneNumber             string `gorm:"index"`
	PhoneNumberID           string
	MessagingProduct        string
	Event                   string
	MessageTemplateID       int64
	MessageTemplateName     string
	MessageTemplateLanguage string
	Field                   string
	Reason                  string
}
