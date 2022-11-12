package entity

import "gorm.io/gorm"

type Role struct {
	gorm.Model  `json:"-"`
	Id          uint8  `gorm:"primary_key" json:"id"`
	Name        string `gorm:"type:varchar(10)" json:"name"`
	Description string `gorm:"type:varchar(50)" json:"description"`
}
