package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model  `json:"-"`
	Id          int64        `json:"id" gorm:"primary_key"`
	Username    string       `json:"username" gorm:"type:varchar(16)"`
	Password    string       `json:"password" gorm:"type:varchar(16)"`
	Sex         string       `json:"sex" gorm:"type:char(1)"`
	NickName    string       `json:"nick_name" gorm:"type:varchar(10)"`
	RoleId      uint8        `json:"role_id"`
	Role        Role         `json:"role" gorm:"foreignKey:RoleId;references:Id"`
	Permissions []Permission `gorm:"many2many:user_permission"`
}

/**
func NewUser(id int64,username string,password string,sex string,nickName string,role string)User{
	return User{id,username,password,sex,nickName,role}
}

func (user User) SetId(id int64){
	user.id=id
}

func (user User) SetUsername(username string){
	user.username=username
}

*/
