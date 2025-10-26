package models

import (
	"encoding/json"
	"time"
)

// WhatsAppMessageDB represents the main message storage in database
type WhatsAppMessageDB struct {
	BaseModel
	// Message identifiers
	MessageID   string    `gorm:"uniqueIndex;type:varchar(100);not null" json:"message_id"`
	Timestamp   time.Time `gorm:"index" json:"timestamp"`
	MessageType string    `gorm:"type:varchar(20)" json:"type"`

	// Message content fields
	TextContent   *string `gorm:"type:text" json:"text_content"`
	MediaType     string  `gorm:"type:varchar(20)" json:"media_type"`
	MediaID       string  `gorm:"type:varchar(100)" json:"media_id"`
	MediaMimeType string  `gorm:"type:varchar(50)" json:"media_mime_type"`
	MediaSha256   string  `gorm:"type:varchar(64)" json:"media_sha256"`
	MediaFilename string  `gorm:"type:varchar(255)" json:"media_filename"` // Original filename for documents
	IsAnimated    bool    `gorm:"default:false" json:"is_animated"`
	StorageURL    string  `gorm:"type:text" json:"storage_url"` // Add this new field

	// Contact reference (foreign key)
	WaId      string          `gorm:"index;type:varchar(20);not null" json:"wa_id"`
	ContactID uint            `gorm:"index" json:"contact_id"`
	Contact   WhatsAppContact `gorm:"foreignKey:ContactID" json:"-"`

	// Template information
	TemplateID       int64  `gorm:"default:0" json:"template_id"`
	TemplateName     string `gorm:"type:varchar(100)" json:"template_name"`
	TemplateLanguage string `gorm:"type:varchar(10)" json:"template_language"`

	// Additional metadata
	RawPayload json.RawMessage `gorm:"type:jsonb" json:"raw_payload"`
	Status     string          `gorm:"type:varchar(20);default:'received'" json:"status"`
}

// Add this TableName method to ensure consistent table naming
func (WhatsAppMessageDB) TableName() string {
	return "whats_app_messages" // Note the underscores, matching naming convention
}
