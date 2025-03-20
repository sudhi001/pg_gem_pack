package models

import "time"

// EmployeePresenceTracker stores details about an employee's presence and movements.
type EmployeePresenceTracker struct {
	ID               uint          `gorm:"primaryKey;autoIncrement"`        // Auto-incremented unique ID
	OrgUserID        int           `gorm:"not null"`                        // Organization User ID (Reference to user)
	OrgID            int           `gorm:"not null"`                        // Organization ID (Reference to organization)
	EmployeeCode     string        `gorm:"type:varchar(50);not null;index"` // Unique employee identifier
	IsInside         bool          `gorm:"not null;default:false"`          // Whether the employee is inside the organization
	LastEntryTime    time.Time     `gorm:"default:null"`                    // Last recorded entry time
	LastExitTime     time.Time     `gorm:"default:null"`                    // Last recorded exit time
	TotalEntries     int           `gorm:"not null;default:0"`              // Total times the employee entered
	TotalExits       int           `gorm:"not null;default:0"`              // Total times the employee exited
	TotalMovements   int           `gorm:"not null;default:0"`              // Total times the employee entered and exited (combined)
	TotalTimeInside  time.Duration `gorm:"not null;default:0"`              // Total time spent inside the organization
	TotalTimeOutside time.Duration `gorm:"not null;default:0"`              // Total time spent outside the organization
	CreatedAt        time.Time     `gorm:"autoCreateTime"`                  // Record creation timestamp
	UpdatedAt        time.Time     `gorm:"autoUpdateTime"`                  // Record last updated timestamp
}
