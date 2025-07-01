package models

// User represents the User table in the database.
type UserDefaultBillingAddress struct {
	BaseModel
	UserID              string  `gorm:"uniqueIndex;not null"` // Unique identifier for the user
	FirstName           *string `gorm:"default:null"`
	LastName            *string `gorm:"default:null"`
	BuildingName        *string `gorm:"default:null"`
	BuildingNumber      *string `gorm:"default:null"`
	LocationName        *string `gorm:"default:null"`
	PinCode             *string `gorm:"default:null"`
	City                *string `gorm:"default:null"`
	State               *string `gorm:"default:null"`
	Country             *string `gorm:"default:null"`
	UpdatedFromDeviceID *string `gorm:"default:null"`
}
