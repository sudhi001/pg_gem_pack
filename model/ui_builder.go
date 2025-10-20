package models

import (
	"time"
)

// Import JSON type from base.go
// type JSON map[string]interface{}

// UIBuildConfig represents a top-level UI builder config (e.g., for a screen)
type UIBuildConfig struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"size:100;index" json:"name"`                   // e.g., HOME_CONTENT
	Type      string    `gorm:"size:50;index" json:"type"`                    // e.g., DYNAMIC
	Meta      *JSON     `gorm:"type:json;default:null" json:"meta,omitempty"` // For future extensibility
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Views     []UIView  `gorm:"foreignKey:ConfigID" json:"views"`
}

// UIView represents a single view/component in the UI builder (can be nested)
type UIView struct {
	ID            uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	ConfigID      uint      `gorm:"index" json:"configId"`
	ParentViewID  *uint     `gorm:"index;default:null" json:"parentViewId,omitempty"` // For nesting
	Order         int       `gorm:"index" json:"order"`                               // For ordering views
	Type          string    `gorm:"size:100;index" json:"type"`                       // e.g., DYNAMIC_PROPERTY_VIEW, TYPE_TWO_CARD, etc.
	Subkey        *string   `gorm:"size:100;default:null" json:"subkey,omitempty"`
	Title         *string   `gorm:"size:255;default:null" json:"title,omitempty"`
	Body          *string   `gorm:"type:text;default:null" json:"body,omitempty"`
	Roles         *JSON     `gorm:"type:json;default:null" json:"roles,omitempty"`         // Array of roles
	ExcludedRoles *JSON     `gorm:"type:json;default:null" json:"excludedRoles,omitempty"` // Array of roles
	Instructions  *JSON     `gorm:"type:json;default:null" json:"instructions,omitempty"`  // Array of instructions
	List          *JSON     `gorm:"type:json;default:null" json:"list,omitempty"`          // Array of items (for cards, grids, etc.)
	Actions       *JSON     `gorm:"type:json;default:null" json:"actions,omitempty"`       // Array of actions (for buttons, etc.)
	Extra                      *JSON     `gorm:"type:json;default:null" json:"extra,omitempty"`                    // For future extensibility
	Enabled                    bool      `gorm:"default:true;index" json:"enabled"`                                 // Flag to enable/disable view rendering
	OnlyForAuthenticatedUser   bool      `gorm:"default:false;index" json:"onlyForAuthenticatedUser"`               // Flag to require authentication for view
	BuildNumber                *int      `gorm:"default:null" json:"buildNumber,omitempty"`                         // Build number for version filtering
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
	Children      []UIView  `gorm:"foreignKey:ParentViewID" json:"children,omitempty"`
}

// TableName for UIBuildConfig
func (UIBuildConfig) TableName() string {
	return "ui_build_configs"
}

// TableName for UIView
func (UIView) TableName() string {
	return "ui_views"
}
