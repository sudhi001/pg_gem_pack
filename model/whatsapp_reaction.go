package models

import (
	"time"
)

// WhatsAppReaction represents a reaction to a WhatsApp message
type WhatsAppReaction struct {
	BaseModel
	MessageID        string    `gorm:"type:varchar(100);not null;index:idx_reaction_message" json:"message_id"`
	ReactedMessageID string    `gorm:"type:varchar(100);not null;index" json:"reacted_message_id"`
	WaId             string    `gorm:"index;type:varchar(20);not null" json:"wa_id"`
	ContactID        uint      `gorm:"index" json:"contact_id"`
	Emoji            string    `gorm:"type:varchar(50)" json:"emoji"`
	Timestamp        time.Time `gorm:"index" json:"timestamp"`
	
	Contact          WhatsAppContact `gorm:"foreignKey:ContactID" json:"-"`
}

func (WhatsAppReaction) TableName() string {
	return "whats_app_reactions"
}

