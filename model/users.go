package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

// User represents the User table in the database.
type User struct {
	BaseModel
	UserID                string           `gorm:"uniqueIndex;not null"` // Unique identifier for the user
	ExternalUserID        *string          `gorm:"index"`                // Nullable external user ID
	Email                 *string          `gorm:"index"`                // Nullable email address
	Password              *string          `gorm:"default:null"`         // Nullable password
	PhoneNumber           *string          `gorm:"index"`                // Phone number
	PhoneNumberExtension  *string          `gorm:"default:null"`         // Phone number extension
	Name                  *string          `gorm:"index"`                // Nullable user name
	ProfileImage          *string          `gorm:"default:null"`         // Nullable profile image URL
	EmailVerified         bool             `gorm:"default:false"`        // Email verification status
	EmailValidationID     *string          `gorm:"default:null"`         // Nullable email validation ID
	PhoneNumberVerified   bool             `gorm:"default:false"`        // Phone number verification status
	Blocked               bool             `gorm:"default:false"`        // Block status
	Activated             bool             `gorm:"default:false"`        // Activation status
	KYCCompleted          bool             `gorm:"default:false"`        // KYC completion status
	RequestHash           *string          `gorm:"default:null"`         // Nullable request hash
	Permissions           PermissionString `gorm:"type:text[];default:NULL"`
	ReasonForBlock        *string          `gorm:"default:null"` // Nullable reason for block
	ReasonForDeactivation *string          `gorm:"default:null"` // Nullable reason for deactivation
	AppCustomerCode       *int             `gorm:"default:null"` // ACME customer code
}

// Strings is a custom type to handle []string serialization/deserialization.
type PermissionString []string

// Scan implements the sql.Scanner interface for PermissionString.
func (s *PermissionString) Scan(value interface{}) error {
	if value == nil {
		*s = []string{}
		return nil
	}

	switch v := value.(type) {
	case []byte:
		// Attempt to unmarshal as JSON
		if err := json.Unmarshal(v, s); err != nil {
			// If unmarshalling fails, try treating it as a comma-separated string
			*s = strings.Split(string(v), ",")
		}
	case string:
		// Attempt to unmarshal as JSON
		if err := json.Unmarshal([]byte(v), s); err != nil {
			// If unmarshalling fails, try treating it as a comma-separated string
			*s = strings.Split(v, ",")
		}
	default:
		return fmt.Errorf("failed to scan PermissionString: expected []byte or string, got %T", value)
	}

	return nil
}

// Value implements the driver.Valuer interface for Strings.
func (s PermissionString) Value() (driver.Value, error) {
	return json.Marshal(s)
}
