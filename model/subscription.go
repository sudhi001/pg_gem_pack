package models

type Subscription struct {
	BaseModel
	UID     string `gorm:"index"`
	UserID  string `gorm:"index;not null"`
	Topic   string `gorm:"index;not null"`
	Enabled bool   `gorm:"default:true;not null"`
	SpaceUID string `gorm:"index;not null"`
}

func (Subscription) TableName() string {
	return "subscriptions"
}

