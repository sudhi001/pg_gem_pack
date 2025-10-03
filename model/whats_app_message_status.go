package models

import "time"

// Create WWhatsAppMessageStatus struct
// WhatsAppMessageStatus represents the status of a WhatsApp message
type WhatsAppMessageStatus struct {
	BaseModel
	MessageID               string    `gorm:"uniqueIndex:idx_message_id;not null;type:varchar(255)" json:"message_id"`
	RecipientID             string    `gorm:"index:idx_recipient;type:varchar(20)" json:"recipient_id"`
	Status                  string    `gorm:"type:varchar(50);not null" json:"status"`
	Timestamp               time.Time `gorm:"not null;index:idx_timestamp" json:"timestamp"`
	PhoneNumber             string    `gorm:"index:idx_phone;type:varchar(20)" json:"phone_number"`
	PhoneNumberID           string    `gorm:"type:varchar(50)" json:"phone_number_id"`
	MessagingProduct        string    `gorm:"type:varchar(50);default:'whatsapp'" json:"messaging_product"`
	Event                   string    `gorm:"type:varchar(50)" json:"event"`
	MessageTemplateID       int64     `gorm:"default:0" json:"message_template_id"`
	MessageTemplateName     string    `gorm:"type:varchar(100)" json:"message_template_name"`
	MessageTemplateLanguage string    `gorm:"type:varchar(10)" json:"message_template_language"`
	Field                   string    `gorm:"type:varchar(50)" json:"field"`
	Reason                  string    `gorm:"type:varchar(255)" json:"reason"`
}
