package models

import "time"

type User struct {
	ID         uint         `gorm:"primaryKey" json:"id"`
	Email      string       `json:"email" gorm:"not null;unique;default:null" validate:"required"`
	Password   string       `json:"password" gorm:"not null;default:null" validate:"required"`
	UserName   string       `json:"userName" gorm:"not null;unique;default:null" validate:"required"`
	Attendance []Attendance `gorm:"foreignKey:UserId"`
	CreatedAt  time.Time    `json:"createdAt"`
}
