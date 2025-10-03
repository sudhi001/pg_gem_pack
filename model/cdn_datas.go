package models

import "encoding/json"

// CDNDatas represents the cdn_datas table in the database
type CDNDatas struct {
	BaseModel
	UID  string          `json:"uid" db:"uid"`
	Name string          `json:"name" db:"name"`
	Data json.RawMessage `json:"data" gorm:"type:json"`
}

type CDNDataHistory struct {
	BaseModel
	UID  string          `json:"uid" db:"uid"`
	Name string          `json:"name" db:"name"`
	Hash string          `json:"hash" db:"hash"`
	Data json.RawMessage `json:"data" gorm:"type:json"`
}
