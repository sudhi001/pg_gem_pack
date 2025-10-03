package models

// FCMProcessingAudit represents audit records for FCM message processing
type FCMProcessingAudit struct {
	BaseModel
	UID             string  `gorm:"index"`
	MessageKey      string  `json:"message_key" gorm:"type:varchar(255);not null"`
	MessageValue    string  `json:"message_value" gorm:"type:text"`
	ErrorType       string  `json:"error_type" gorm:"type:varchar(100);not null"`
	ErrorMessage    string  `json:"error_message" gorm:"type:text;not null"`
	ProcessingStage string  `json:"processing_stage" gorm:"type:varchar(100);not null"`
	RetryCount      int     `json:"retry_count" gorm:"default:0"`
	IsResolved      bool    `json:"is_resolved" gorm:"default:false"`
	ResolutionNote  *string `json:"resolution_note" gorm:"type:text"`
}

// TableName specifies the table name for FCMProcessingAudit
func (FCMProcessingAudit) TableName() string {
	return "fcm_processing_audit"
}
