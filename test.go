package main

import (
	"fmt"
	"gorm-study/config"
	"gorm-study/entity"
	"gorm-study/service"
	"gorm-study/utils/token"
)

func main() {
	user := entity.User{Id: 2, Username: "root1", Password: "123456", Sex: "男", NickName: "张三"}
	accessToken := token.CreateAccessToken(user)
	println(token.ParseTokenFormToUsername(accessToken))
	//fmt.Println(accessToken)
	//refreshToken := token.CreateRefreshToken(user)
	//fmt.Println(refreshToken)

}

var userService = service.NewUserService()

func test1() {
	config.DB.AutoMigrate(&entity.User{})

	var user = entity.User{Id: 1, Username: "root", Password: "123456", Sex: "男"}

	create := config.DB.Create(&user)

	fmt.Println(create)
}

func test2() {
	users := userService.FindAllUser()
	fmt.Println(users)
}

func test3() {
	var user = userService.FindUserById(1)
	fmt.Println(user)
}

func test4() {
	var user = entity.User{Id: 2, Username: "root2123456", Password: "123456789", Sex: "女", NickName: "王红"}
	result := userService.SaveUser(user)
	fmt.Println(result)
}

func test5() {
	var user = entity.User{Id: 2, Username: "root2123456", Password: "123456789", Sex: "女", NickName: "王红28"}
	updateUser := userService.UpdateUser(user)
	fmt.Println(updateUser)
}

func test6() {
	var user = userService.FindByUsername("root")
	fmt.Println(user)
}
