package models



type SecurityAccessLog struct {
    BaseModel
    OrgID            int       `gorm:"index"`
    OrgUserID        int       `gorm:"index"`
    EventTimeEpoch   int64
    Direction        string    `gorm:"type:varchar(10)"`
    AccessType       string    `gorm:"type:varchar(50)"`
    EventID          int
    AccessPointID    int       `gorm:"index"`
    EmployeeCode     string    `gorm:"type:varchar(50);index"`
    AccessPointName  string    `gorm:"type:varchar(100)"`
    MobileAccessMode string    `gorm:"type:varchar(50)"`
    SiteID          int       `gorm:"index"`
    Name            string    `gorm:"type:varchar(100)"`
}