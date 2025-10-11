package models

import "time"

type GoldMetalRateHistory struct {
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

