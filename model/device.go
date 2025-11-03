package models

import (
	"time"
)

// Device represents a device entity in the database
type Device struct {
	BaseModel
	DeviceID           string    `gorm:"index;not null"`
	PrivateKey         string    `gorm:"not null"`
	PublicKey          string    `gorm:"not null"`
	FCM                string    `gorm:"default:null"`
	APNS               string    `gorm:"default:null"`
	OS                 string    `gorm:"default:null"`
	OSVersion          string    `gorm:"default:null"`
	AppBuild           string    `gorm:"default:null"`
	AppVersion         string    `gorm:"default:null"`
	Brand              string    `gorm:"default:null"`
	Model              string    `gorm:"default:null"`
	ClientIP           string    `gorm:"default:null"`
	UserAgent          string    `gorm:"default:null"`
	LastActiveAt       time.Time `gorm:"autoUpdateTime"`
	LastKnownLocation  *string   `gorm:"type:geometry(Point,4326);default:null"` // SRID 4326 for geographic coordinates
	Locale             string    `gorm:"default:null"`
	Timezone           string    `gorm:"default:null"`
	RequestTime        string    `gorm:"default:null"`
	SelectedBranchCode *string   `gorm:"default:null"`       // User's selected branch code
	DeviceInstanceID   string    `gorm:"index;default:null"` // Client-generated unique device instance ID for deduplication
}
