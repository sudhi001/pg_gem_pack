package models

// DeviceEvent represents an analytics event from a device
type DeviceEvent struct {
	BaseModel
	IP          string  `gorm:"index"`
	DeviceID    *string `gorm:"index"`
	UserID      *string `gorm:"index"`
	EventType   string  `gorm:"index;not null"`
	Description string  `gorm:"type:text"`
	Metadata    JSON    `gorm:"type:jsonb"`
	UserAgent   string  `gorm:"type:text"`
	ClientIP    string  `gorm:"index"`
}

