package models

import "encoding/json"

// Branche model for handling payment details

type MetalRate struct {
	BaseModel
	BaseCode           string            `gorm:"index"`     // Index for fast searching
	ConversionRates    json.RawMessage   `gorm:"type:json"` // GORM will ignore this field for DB mapping (we'll store as JSON)
	TimeLastUpdateUnix int64             `gorm:"index"`     // Unix timestamp for last update
	BranchCode         string            `gorm:"size:255"`
	Date               string            `gorm:"size:10;index"`                                                     // Date in `YYYY-MM-DD` format
	MetalRateDetails   []MetalRateDetail `gorm:"foreignKey:BranchID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // Relation to MetalRateDetail
}

type MetalRateDetail struct {
	BaseModel
	BranchID            uint    `gorm:"not null"` // Foreign key to Branch
	ItemType            string  `gorm:"size:255"` // e.g., Gold, Silver
	Purity              float64 // Purity percentage
	SaleRate            float64 // Sale rate
	Unit                string  `gorm:"size:255"` // Unit of measurement (e.g., per gram)
	Date                string  // Date of the rate
	LastUpdatedDateTime string  // Last update timestamp

}
