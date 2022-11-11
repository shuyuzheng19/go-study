package service

import (
	"encoding/base64"
	"errors"
	"gorm-study/common"
	"gorm-study/dao"
	"gorm-study/dto"
	"gorm-study/entity"
	"gorm-study/utils/redis"
	"gorm-study/utils/token"
	"gorm-study/vo"
)

var userDao = dao.NewUserDao()

type userService struct {
}

var myUserService = userService{}

func NewUserService() userService {
	return myUserService
}

func (*userService) Login(dto dto.UserDto) vo.ResponseToken {

	user := userDao.FindByUsernameAndPassword(dto)

	if user.Id == 0 {
		return vo.ResponseToken{}
	}

	var username = base64.StdEncoding.EncodeToString([]byte(user.Username))

	var accessToken = token.CreateAccessToken(user)

	var refreshToken = token.CreateRefreshToken(user)

	a := redis.Set(common.AccessTokenPrefix+username, accessToken, common.AccessTokenExpireTime)

	b := redis.Set(common.RefreshTokenPrefix+username, refreshToken, common.RefreshTokenExpireTime)

	if a && b {
		return vo.ResponseToken{Username: user.Username, AccessToken: accessToken, RefreshToken: refreshToken}
	} else {
		return vo.ResponseToken{}
	}

}

func (*userService) FindByUsername(username string) entity.User {

	user := userDao.FindByUsername(username)

	if user.Id == 0 {
		panic(errors.New("找不到该用户名"))
	}

	return user
}

func (*userService) FindAllUser() []entity.User {

	users := userDao.FindAll()

	if len(users) == 0 {
		return []entity.User{}
	} else {
		return users
	}

}

func (*userService) SaveUser(user entity.User) bool {
	user.Id = 0

	if user.Username == "" {
		panic(errors.New("用户名称不能为空"))
	}

	var userNameSize = len(user.Username)

	if !(userNameSize < 16 && userNameSize > 8) {
		panic(errors.New("用户名称长度不能大于16并且不能8个"))
	}

	if user.Password == "" {
		panic(errors.New("用户名称不能为空"))
	}

	var passWordSize = len(user.Password)

	if !(passWordSize < 16 && passWordSize > 8) {
		panic(errors.New("用户名称长度不能大于16并且不能8个"))
	}

	if user.Sex != "男" && user.Sex != "女" {
		panic(errors.New("性别只能是男或者女"))
	}

	resultUser := userDao.AddUser(user)

	if resultUser.Id == 0 {
		return false
	}

	return true
}

func (*userService) UpdateUser(user entity.User) bool {

	if user.Id <= 0 {
		panic(errors.New("错误的ID参数"))
	}

	if !userDao.IsExistsUser(user.Id) {
		return false
	}

	userDao.UpdateUser(user)

	return true
}

func (*userService) FindUserById(id int64) entity.User {
	user := userDao.FindById(id)

	if user.Id == 0 {
		panic(errors.New("该用户不存在"))
	}

	return user
}
