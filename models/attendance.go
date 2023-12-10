package models

import "time"

type Attendance struct {
	ID                uint       `gorm:"primaryKey" json:"id"`
	UserId            int        `json:"userId"`
	ClockIn           time.Time  `json:"clockIn"`
	ClockInIpAddress  string     `json:"clockInIpAddress"`
	ClockInLongitude  string     `json:"clockInLongitude"`
	ClockInLatitude   string     `json:"clockInLatitude"`
	ClockOut          *time.Time `json:"clockOut"`
	ClockOutIpAddress string     `json:"clockOutIpAddress"`
	ClockOutLongitude string     `json:"clockOutLongitude"`
	ClockOutLatitude  string     `json:"clockOutLatitude"`
	CreatedAt         time.Time  `json:"createdAt"`
}
