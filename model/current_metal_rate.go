package models

// CurrentMetalRate model for handling current metal rate information
type CurrentMetalRate struct {
	BaseModel
	UID                 string  `gorm:"index"` // Unique identifier, indexed
	ItemType            string  `gorm:"index"` // Type of metal (Gold, Silver, etc.), indexed
	ItemTypeCode        int     `json:"itemTypeCode"`
	Purity              float64 `json:"purity"`
	SaleRate            float64 `json:"saleRate"`
	Unit                string  `json:"unit"`
	Date                string  `json:"date"`
	LastUpdatedDateTime string  `json:"lastUpdatedDateTime"`
	BranchCode          string  `gorm:"index"` // Branch code this rate applies to, indexed
}

func (CurrentMetalRate) TableName() string {
	return "metal_rates"
}

// GoldMetalRate is an alias for backward compatibility
type GoldMetalRate = CurrentMetalRate

