package models

import (
	"time"

	"gorm.io/gorm"
)

type Address struct {
	ID         string `gorm:"size:36;not null;uniqueIndex;primary_key;"`
	User       User
	UserID     string `gorm:"size:36;not null;index;"`
	Name       string `gorm:"size:100;"`
	IsPrimary  bool
	CityID     string `gorm:"size:100;not null;"`
	ProvinceID string `gorm:"size:100;not null;"`
	Address1   string `gorm:"size:255;"`
	Address2   string `gorm:"size:255;"`
	Phone      string `gorm:"size:100;"`
	Email      string `gorm:"size:100;"`
	PostCode   string `gorm:"size:5;"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}
