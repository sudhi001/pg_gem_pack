package models

type AIGeneratedNotification struct {
	BaseModel
	UID          string  `gorm:"uniqueIndex;not null" json:"uid"`
	ItemType     string  `gorm:"size:50;not null" json:"itemType"`
	ItemTypeCode int     `gorm:"not null" json:"itemTypeCode"`
	Purity       float64 `gorm:"type:decimal(5,2);not null" json:"purity"`
	SaleRate     float64 `gorm:"type:decimal(10,2);not null" json:"saleRate"`
	Unit         string  `gorm:"size:20;not null" json:"unit"`
	Date         string  `gorm:"size:50;not null;index" json:"date"`
	BranchCode   string  `gorm:"size:50;not null;index" json:"branchCode"`

	NotificationTitle string `gorm:"size:255;not null" json:"notificationTitle"`
	NotificationBody  string `gorm:"type:text;not null" json:"notificationBody"`
	AIModelUsed       string `gorm:"size:100" json:"aiModelUsed"`
	PromptUsed        string `gorm:"type:text" json:"promptUsed"`
	AIResponseRaw     string `gorm:"type:text" json:"aiResponseRaw"`
	GenerationStatus  string `gorm:"size:50;default:'success';index" json:"generationStatus"`
	ErrorMessage      string `gorm:"type:text" json:"errorMessage"`
}

func (AIGeneratedNotification) TableName() string {
	return "ai_generated_notifications"
}
