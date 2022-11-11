package router

import (
	"github.com/gin-gonic/gin"
	"gorm-study/controller"
)

type projectRouter struct {
	group gin.RouterGroup
}

func NewProjectRouter(group gin.RouterGroup) projectRouter {
	return projectRouter{group: group}
}

func (router *projectRouter) AddUserRouter() {
	group := router.group

	userController := controller.NewUserController()

	routerGroup := group.Group("/api/v1/user")

	{

		routerGroup.GET("/findAll", userController.FindAllUser)

		routerGroup.GET("/id/:id", userController.FindByUserId)

		routerGroup.POST("/login", userController.Login)

	}

}
