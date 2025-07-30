package models

// Branche model for handling payment details

type WhatsAppTemplates struct {
	BaseModel
	MessageTemplateId       string  `json:"message_template_id" gorm:"not null"`         // eg: 857032099634334
	MessageTemplateName     string  `json:"message_template_name" gorm:"index;not null"` // Index for fast searching
	MessageTemplateLanguage string  `json:"message_template_language" gorm:"size:5"`     // eg: ml, en, etc.
	Status                  string  `json:"status" gorm:"size:20;not null"`              // eg: APPROVED
	ImageURL                *string `json:"image_url" gorm:"default:null"`               // URL for template image
}
