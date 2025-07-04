package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

// BaseModel holds common fields for all models
type BaseModel struct {
	ID           uint      `gorm:"primaryKey;autoIncrement;index"`
	SpaceUID     string    `gorm:"size:50;null"`
	CreatedByUID string    `gorm:"index;null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
	Active       bool      `gorm:"default:true"` // Flag to identify if this is active or disabled
}

// JSON type for handling JSON data in Go
type JSON map[string]interface{}

// Scan implements the sql.Scanner interface for JSON
func (j *JSON) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}

	switch v := value.(type) {
	case []byte:
		return json.Unmarshal(v, j)
	case string:
		return json.Unmarshal([]byte(v), j)
	default:
		return fmt.Errorf("cannot scan %T into JSON", value)
	}
}

// Value implements the driver.Valuer interface for JSON
func (j JSON) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}
