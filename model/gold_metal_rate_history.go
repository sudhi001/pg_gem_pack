package models

import "time"

type MetalRateHistory struct {
	BaseModel
	UID                 string    `gorm:"index"`
	ItemType            string    `gorm:"index"`
	ItemTypeCode        int       `json:"itemTypeCode"`
	Purity              float64   `gorm:"index" json:"purity"`
	SaleRate            float64   `json:"saleRate"`
	Unit                string    `json:"unit"`
	Date                string    `gorm:"index" json:"date"`
	LastUpdatedDateTime string    `json:"lastUpdatedDateTime"`
	BranchCode          string    `gorm:"index" json:"branchCode"`
	SnapshotDate        time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"snapshotDate"`
}

func (MetalRateHistory) TableName() string {
	return "metal_rates_history"
}

// GoldMetalRateHistory is an alias for backward compatibility
type GoldMetalRateHistory = MetalRateHistory

