package models

// OrgDeviceMapperMapper stores user access details
type OrgDeviceMapperMapper struct {
	ID           uint   `gorm:"primaryKey;autoIncrement"` // Auto-incremented ID
	OrgID        int    `gorm:"not null"`                 // Organization ID
	AccessPointName string `gorm:"type:varchar(100);not null"`
	Direction string `gorm:"type:varchar(50);not null"`
	AccessType string `gorm:"type:varchar(50);not null"`
}
