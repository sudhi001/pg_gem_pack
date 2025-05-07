package models

import "time"

// EmployeeAttendanceReport represents aggregated attendance data for reporting
type EmployeeAttendanceReport struct {
	BaseModel
	EmployeeCode   string    `gorm:"index"`           // Employee code, indexed
	StartDate      time.Time `gorm:"index;type:date"` // Start date of the report period
	EndDate        time.Time `gorm:"index;type:date"` // End date of the report period
	TotalDays      int       `gorm:"default:0"`       // Total number of days in the period
	WeekdayCount   int       `gorm:"default:0"`       // Number of weekdays in the period (excluding weekends)
	WeekendCount   int       `gorm:"default:0"`       // Number of weekend days in the period
	PresentDays    int       `gorm:"default:0"`       // Number of days present
	AbsentDays     int       `gorm:"default:0"`       // Number of days absent
	HalfDays       int       `gorm:"default:0"`       // Number of half days
	WeekendPresent int       `gorm:"default:0"`       // Number of weekend days when employee was present

	// Time metrics
	TotalTimeIn    int       `gorm:"default:0"`    // Total time inside in minutes for the period
	AvgDailyTimeIn float64   `gorm:"default:0"`    // Average daily time in minutes
	WeekdayAvgTime float64   `gorm:"default:0"`    // Average time for weekdays only
	WeekendAvgTime float64   `gorm:"default:0"`    // Average time for weekend days only
	StandardHours  int       `gorm:"default:480"`  // Standard work hours in minutes (default 8 hours = 480 minutes)
	ExtraTimeTotal int       `gorm:"default:0"`    // Total extra time worked in minutes beyond standard hours
	ShortTimeTotal int       `gorm:"default:0"`    // Total time short of standard hours
	DaysWithExtra  int       `gorm:"default:0"`    // Number of days with extra time worked
	DaysShort      int       `gorm:"default:0"`    // Number of days with less than standard hours
	MaxExtraTime   int       `gorm:"default:0"`    // Maximum extra time worked in a single day (minutes)
	MaxExtraDate   time.Time `gorm:"default:null"` // Date when maximum extra time was worked

	// Time pattern metrics
	EarliestCheckin  time.Time `gorm:"default:null"`     // Earliest check-in during the period
	LatestCheckin    time.Time `gorm:"default:null"`     // Latest check-in during the period
	EarliestCheckout time.Time `gorm:"default:null"`     // Earliest check-out during the period
	LatestCheckout   time.Time `gorm:"default:null"`     // Latest check-out during the period
	AvgCheckinTime   string    `gorm:"type:varchar(5)"`  // Average check-in time (HH:MM format)
	AvgCheckoutTime  string    `gorm:"type:varchar(5)"`  // Average check-out time (HH:MM format)
	MostFrequentDay  string    `gorm:"type:varchar(10)"` // Most frequent day of week present

	// Streak metrics
	LongestStreak      int       `gorm:"default:0"`    // Longest streak of consecutive present days
	LongestStreakStart time.Time `gorm:"default:null"` // Start date of longest streak
	LongestStreakEnd   time.Time `gorm:"default:null"` // End date of longest streak
	CurrentStreak      int       `gorm:"default:0"`    // Current streak of consecutive present days
	CurrentStreakStart time.Time `gorm:"default:null"` // Start date of current streak

	// Pattern metrics
	LateArrivals    int `gorm:"default:0"` // Number of late arrivals (after 9:00 AM)
	EarlyDepartures int `gorm:"default:0"` // Number of early departures (before 5:00 PM)
	PerfectDays     int `gorm:"default:0"` // Days with on-time arrival and departure

	// Access metrics
	AccessPoints   string  `gorm:"type:text"`        // JSON list of access points used
	MostUsedAccess string  `gorm:"type:varchar(50)"` // Most frequently used access point
	AccessCount    int     `gorm:"default:0"`        // Total number of access events
	AvgDailyAccess float64 `gorm:"default:0"`        // Average number of access events per day

	// Report metadata
	ReportType string `gorm:"type:varchar(20);index"`  // Type of report (Daily, Weekly, Monthly, Custom)
	ReportHash string `gorm:"type:varchar(64);unique"` // Hash to prevent duplicate reports
}
