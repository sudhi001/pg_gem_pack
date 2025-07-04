package models

import (
	"time"
)

// ExternalServiceAudit tracks external service API failures
type ExternalServiceAudit struct {
	UID             string     `json:"uid" gorm:"primaryKey;type:varchar(36)"`
	ServiceName     string     `json:"service_name" gorm:"not null;index"`
	Endpoint        string     `json:"endpoint" gorm:"not null"`
	Method          string     `json:"method" gorm:"not null"`
	RequestData     string     `json:"request_data" gorm:"type:text"`
	ResponseData    string     `json:"response_data" gorm:"type:text"`
	StatusCode      int        `json:"status_code"`
	ErrorMessage    string     `json:"error_message" gorm:"type:text"`
	RetryCount      int        `json:"retry_count" gorm:"default:0"`
	IsResolved      bool       `json:"is_resolved" gorm:"default:false"`
	ResolvedAt      *time.Time `json:"resolved_at"`
	ResolvedBy      string     `json:"resolved_by"`
	ResolutionNotes string     `json:"resolution_notes" gorm:"type:text"`
	BaseModel
}

// TableName specifies the table name for ExternalServiceAudit
func (ExternalServiceAudit) TableName() string {
	return "external_service_audits"
}
