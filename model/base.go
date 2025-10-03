package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

// BaseModel holds common fields for all models
type BaseModel struct {
	ID           uint      `json:"uid" gorm:"primaryKey;autoIncrement;index"`
	SpaceUID     string    `json:"space_uid" gorm:"size:50;null"`
	CreatedByUID string    `json:"created_by_uid" gorm:"index;null"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Active       bool      `json:"active" gorm:"default:true"` // Flag to identify if this is active or disabled
}

// JSON type for handling JSON data in Go (supports both objects and arrays)
type JSON struct {
	Data interface{}
}

// JSONPtr is a pointer to JSON for database operations
type JSONPtr struct {
	JSON JSON
}

// Original JSONObject type for structured objects only
type JSONObject map[string]interface{}

// Scan implements the sql.Scanner interface for JSON
func (j *JSON) Scan(value interface{}) error {
	if value == nil {
		j.Data = nil
		return nil
	}

	switch v := value.(type) {
	case []byte:
		// Try to unmarshal as interface{} to handle both objects and arrays
		var data interface{}
		if err := json.Unmarshal(v, &data); err != nil {
			return err
		}
		j.Data = data
		return nil
	case string:
		// Try to unmarshal as interface{} to handle both objects and arrays
		var data interface{}
		if err := json.Unmarshal([]byte(v), &data); err != nil {
			return err
		}
		j.Data = data
		return nil
	default:
		return fmt.Errorf("cannot scan %T into JSON", value)
	}
}

// Value implements the driver.Valuer interface for JSON
func (j JSON) Value() (driver.Value, error) {
	if j.Data == nil {
		return nil, nil
	}
	return json.Marshal(j.Data)
}

// MarshalJSON implements the json.Marshaler interface for JSON pointer
func (j *JSON) MarshalJSON() ([]byte, error) {
	if j == nil || j.Data == nil {
		return json.Marshal(nil)
	}
	return json.Marshal(j.Data)
}
