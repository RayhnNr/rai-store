package models

import "time"

type Address struct {
	ID         string `gorm:"size:36;not null;primaryKey"`
	User       User
	UserID     string `gorm:"size:36;index"`
	Name       string `gorm:"size:100"`
	IsPrimary  bool
	ProvinceID string `gorm:"size:100"`
	CityID     string `gorm:"size:100"`
	Address    string `gorm:"size:255"`
	Address2   string `gorm:"size:255"`
	Phone      string `gorm:"size:100"`
	Email      string `gorm:"size:100"`
	PostCode   string `gorm:"size:100"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}