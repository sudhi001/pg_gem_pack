package models

// Branche model for handling payment details

type WhatsAppTemplates struct {
	BaseModel
	MessageTemplateId       string `gorm:"not null"`         // eg: 857032099634334
	MessageTemplateName     string `gorm:"index;not null"`   // Index for fast searching
	MessageTemplateLanguage string `gorm:"size:5"`           // eg: ml, en, etc.
	Status                  string `gorm:"size:20;not null"` // eg: APPROVED
}
