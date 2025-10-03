package models

// EmployeeCodeMapper stores user access details
type EmployeeCodeMapper struct {
	ID           uint   `gorm:"primaryKey;autoIncrement"` // Auto-incremented ID
	OrgUserID    int    `gorm:"not null"`                 // Organization User ID
	OrgID        int    `gorm:"not null"`                 // Organization ID
	EmployeeName string `gorm:"type:varchar(100);not null"`
	EmployeeCode string `gorm:"type:varchar(50);not null"`
}
