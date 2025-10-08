package models

import (
	"time"
)

// DeletedUser represents archived user data for reference
type DeletedUser struct {
	BaseModel
	OriginalUserID        string           `gorm:"index;not null"`        // Original user ID before deletion
	ExternalUserID        *string          `gorm:"index"`                  // Nullable external user ID
	Email                 *string          `gorm:"index"`                  // Nullable email address
	Password              *string          `gorm:"default:null"`           // Nullable password
	PhoneNumber           *string          `gorm:"index"`                  // Phone number
	PhoneNumberExtension  *string          `gorm:"default:null"`           // Phone number extension
	Name                  *string          `gorm:"index"`                  // Nullable user name
	ProfileImage          *string          `gorm:"default:null"`           // Nullable profile image URL
	EmailVerified         bool             `gorm:"default:false"`         // Email verification status
	EmailValidationID     *string          `gorm:"default:null"`            // Nullable email validation ID
	PhoneNumberVerified    bool             `gorm:"default:false"`          // Phone number verification status
	Blocked               bool             `gorm:"default:false"`          // Block status
	Activated             bool             `gorm:"default:false"`          // Activation status
	KYCCompleted          bool             `gorm:"default:false"`           // KYC completion status
	RequestHash           *string          `gorm:"default:null"`           // Nullable request hash
	Permissions           *string          `gorm:"type:text;default:NULL"` // Store as JSON string
	ReasonForBlock        *string          `gorm:"default:null"`           // Nullable reason for block
	ReasonForDeactivation *string          `gorm:"default:null"`           // Nullable reason for deactivation
	AppCustomerCode       *int             `gorm:"default:null"`           // ACME customer code
	DeletedAt             time.Time        `gorm:"not null"`              // When the user was deleted
	DeletedBy             string           `gorm:"not null"`              // Who deleted the user (user ID)
	DeletionReason        *string          `gorm:"default:null"`           // Reason for deletion
	OriginalCreatedAt     time.Time        `gorm:"not null"`              // Original creation time
	OriginalUpdatedAt     time.Time        `gorm:"not null"`              // Original last update time
}
