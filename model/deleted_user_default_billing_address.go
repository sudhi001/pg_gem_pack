package models

import (
	"time"
)

// DeletedUserDefaultBillingAddress represents archived billing address data for reference
type DeletedUserDefaultBillingAddress struct {
	BaseModel
	OriginalUserID      string    `gorm:"index;not null"` // Original user ID before deletion
	FirstName           *string   `gorm:"default:null"`
	LastName            *string   `gorm:"default:null"`
	BuildingName        *string   `gorm:"default:null"`
	BuildingNumber      *string   `gorm:"default:null"`
	LocationName        *string   `gorm:"default:null"`
	PinCode             *string   `gorm:"default:null"`
	City                *string   `gorm:"default:null"`
	State               *string   `gorm:"default:null"`
	Country             *string   `gorm:"default:null"`
	UpdatedFromDeviceID *string   `gorm:"default:null"`
	DeletedAt           time.Time `gorm:"not null"` // When the billing address was deleted
	DeletedBy           string    `gorm:"not null"` // Who deleted the billing address (user ID)
	OriginalCreatedAt   time.Time `gorm:"not null"` // Original creation time
	OriginalUpdatedAt   time.Time `gorm:"not null"` // Original last update time
}
