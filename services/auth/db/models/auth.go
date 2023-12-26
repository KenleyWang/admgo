package models

import "gorm.io/gorm"

type Auth struct {
	gorm.Model
	A string `gorm:"size:32" json:"a"`
}
