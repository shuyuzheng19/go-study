package main

import (
	"github.com/gin-gonic/gin"
	"gorm-study/common"
	"gorm-study/filter"
	"gorm-study/router"
	"net/http"
	"runtime/debug"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Next()
	}
}

func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			debug.PrintStack()
			c.JSON(500, common.BuildError("服务器异常"))
			c.Abort()
		}
	}()
	c.Next()
}

func main() {
	server := gin.Default()
	server.Use(Cors(), Recover, filter.GlobalFilter())
	group := server.RouterGroup
	projectRouter := router.NewProjectRouter(group)
	projectRouter.AddUserRouter()
	server.Run(":8089")
}
