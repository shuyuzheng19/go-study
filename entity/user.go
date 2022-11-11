package entity

type User struct {
	Id       int64  `json:"id" gorm:"primary_key"`
	Username string `json:"username" gorm:"type:varchar(16)"`
	Password string `json:"password" gorm:"type:varchar(16)"`
	Sex      string `json:"sex" gorm:"type:char(1)"`
	NickName string `json:"nick_name" gorm:"type:varchar(10)"`
	Role     string `json:"role"`
}
