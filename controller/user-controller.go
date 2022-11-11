package controller

import (
	"github.com/gin-gonic/gin"
	"gorm-study/common"
	"gorm-study/dto"
	"gorm-study/service"
	"net/http"
	"strconv"
)

var userService = service.NewUserService()

type userController struct {
}

func NewUserController() userController {
	return userController{}
}

func (*userController) FindAllUser(ctx *gin.Context) {

	users := userService.FindAllUser()

	var size = len(users)

	if size == 0 {
		ctx.JSON(http.StatusOK, common.BuildFailure("找不到用户"))
	}

	ctx.JSON(http.StatusOK, common.BuildSuccess(users))

}

func (*userController) Login(ctx *gin.Context) {
	var userDto dto.UserDto

	ctx.ShouldBind(&userDto)

	tokenResponse := userService.Login(userDto)

	if tokenResponse.Username == "" {
		ctx.JSON(http.StatusOK, common.BuildFailure("账号或者密码错误!"))
		return
	}

	ctx.JSON(http.StatusOK, common.BuildSuccess(tokenResponse))

}

func (*userController) FindByUserId(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

	if id <= 0 {
		ctx.JSON(http.StatusOK, common.BuildFailure("不合法的ID"))
		return
	}

	var user = userService.FindUserById(int64(id))

	if user.Id == 0 {
		ctx.JSON(http.StatusOK, common.BuildFailure("找不到该用户名"))
		return
	}

	ctx.JSON(http.StatusOK, common.BuildSuccess(user))

}
