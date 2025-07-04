package models

// Transaction model for handling payment details
type Transaction struct {
	// Embed BaseModel to inherit common fields (CreatedByUID, CreatedAt, UpdatedAt)
	BaseModel
	DeviceID        string           `gorm:"size:50;not null"`
	TransactionUID  string           `gorm:"size:255;unique;not null"`
	Amount          float64          `gorm:"not null"`
	Currency        string           `gorm:"size:3;not null"`
	ProductCode     string           `gorm:"size:50"`
	BranchCode      string           `gorm:"size:50"`
	BranchUID       string           `gorm:"size:255"`
	UID             string           `gorm:"size:255"`
	Payload         JSON             `gorm:"type:jsonb;not null"` // Change Payload to use jsonb type in PostgreSQL
	TransactionLogs []TransactionLog `gorm:"foreignKey:TransactionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// TransactionLog model for storing transaction log entries
type TransactionLog struct {
	BaseModel
	TransactionID uint   `gorm:"not null"` // Foreign key referencing Transaction.ID
	Payload       JSON   `gorm:"type:jsonb;not null"`
	Status        string `gorm:"size:50;not null"` // E.g., "INFO", "ERROR", "SUCCESS"
}
