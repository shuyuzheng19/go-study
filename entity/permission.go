package entity

type Permission struct {
	Id          int64  `gorm:"primary_key"`
	Name        string `gorm:"type:varchar(15)"`
	Description string `gorm:"type:varchar(50)"`
}
