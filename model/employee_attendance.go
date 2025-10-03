package models

import "time"

type EmployeeAttendance struct {
	BaseModel
	UID          string    `gorm:"index"`                  // Unique identifier, indexed
	EmployeeCode string    `gorm:"index"`                  // Employee code from security access, indexed
	Date         time.Time `gorm:"index;type:date"`        // Date of attendance (without time component), indexed
	FirstCheckIn time.Time `gorm:"index"`                  // First check-in time of the day, indexed
	LastCheckOut time.Time `gorm:"index;default:null"`     // Last check-out time of the day, indexed
	TotalTimeIn  int       `gorm:"default:0"`              // Total time inside in minutes
	Status       string    `gorm:"type:varchar(20);index"` // Status (Present, Absent, Half-Day)
	Notes        string    `gorm:"type:text"`              // Any additional notes
}
