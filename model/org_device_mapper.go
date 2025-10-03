package models

// OrgDeviceMapper stores user access details
type OrgDeviceMapper struct {
	ID              uint   `gorm:"primaryKey;autoIncrement"` // Auto-incremented ID
	OrgID           int    `gorm:"not null"`                 // Organization ID
	AccessPointName string `gorm:"type:varchar(100);not null"`
	AccessPointId   int    `gorm:"not null"`
	AccessType      string `gorm:"type:varchar(50);not null"`
	IS_GATE         bool   `gorm:"not null;default:false"` // Default value set to false
}
