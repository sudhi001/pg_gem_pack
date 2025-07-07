package models

import (
	"time"
)

// Offer represents the Offers table in the database.
type Offer struct {
	BaseModel
	UID              string     `gorm:"uniqueIndex;not null"`                     // Unique identifier for the offer
	BranchUID        string     `gorm:"index;not null"`                           // Branch UID to which the offer belongs
	Title            string     `gorm:"size:255;not null;index"`                  // Title of the offer
	ShortDescription string     `gorm:"size:500;not null"`                        // Short description of the offer
	Description      string     `gorm:"type:text;not null"`                       // Full description of the offer
	ExpireOn         *time.Time `gorm:"index"`                                    // Expiration date/time of the offer (nullable)
	ThumbnailURL     *string    `gorm:"size:500;default:null"`                    // URL to the thumbnail image (nullable)
	ImageURL         *string    `gorm:"size:500;default:null"`                    // URL to the main image (nullable)
	VideoURL         *string    `gorm:"size:500;default:null"`                    // URL to the video (nullable)
	Type             string     `gorm:"size:50;not null;index;default:'general'"` // Type of offer (e.g., 'discount', 'promotion', 'general', etc.)
	Active           bool       `gorm:"default:true;index"`                       // Whether the offer is active
	Tags             *JSON      `gorm:"type:json;default:null"`                   // Tags associated with the offer (nullable)
	Link             *string    `gorm:"size:500;default:null"`                    // Link URL (can be https:// or relative path like /offer/details)
}

// TableName specifies the table name for the Offer model
func (Offer) TableName() string {
	return "offers"
}
