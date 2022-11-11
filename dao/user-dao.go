package dao

import (
	"gorm-study/config"
	"gorm-study/dto"
	"gorm-study/entity"
)

type userDao struct {
}

func NewUserDao() userDao {
	return userDao{}
}

func (*userDao) FindByUsernameAndPassword(userDto dto.UserDto) entity.User {
	var user entity.User

	config.DB.Where("username=? and password=?", userDto.Username, userDto.Password).Find(&user)

	return user
}

func (*userDao) FindById(id int64) entity.User {
	var user entity.User
	config.DB.First(&user, "id=?", id)
	return user
}

func (*userDao) AddUser(user entity.User) entity.User {
	result := config.DB.Create(&user)
	if err := result.Error; err != nil {
		return user
	} else {
		return entity.User{}
	}
}

func (*userDao) IsExistsUser(id int64) bool {

	result := config.DB.First(&entity.User{}, "id=?", id)

	if err := result.Error; err != nil {
		return false
	}

	return true
}

func (*userDao) FindByUsername(username string) entity.User {

	var user entity.User

	result := config.DB.First(&user, "username=?", username)

	if result.Error != nil {
		return entity.User{}
	}

	return user

}

func (*userDao) UpdateUser(user entity.User) entity.User {
	result := config.DB.Updates(&user)

	if err := result.Error; err != nil {
		return user
	} else {
		return entity.User{}
	}

}

func (*userDao) FindAll() []entity.User {
	var users []entity.User
	config.DB.Find(&users)
	return users
}
