package main

import (
	"fmt"
	"gorm-study/config"
	"gorm-study/entity"
	"gorm-study/service"
)

func main() {
	//config.DB.AutoMigrate(&entity.User{})
	//
	//user := userService.FindUserById(1)
	//
	//fmt.Println(user)

	var map1, map2 = make(map[string]int), make(map[string]int)

	map1["a"] = 1
	map2["v"] = 2

	for key, _ := range map1 {
		i, ok := map2[key]
		fmt.Println(i, ok)
	}

	//
	//var user=entity.User{
	//	Id:       1,
	//	Username: "z2528959216",
	//	Password: "123456789",
	//	Sex:      "男",
	//	NickName: "郑书宇",
	//	RoleId:   1,
	//}
	//
	//saveUser := userService.SaveUser(user)
	////
	//fmt.Println(saveUser)

	//config.DB.AutoMigrate(&entity.Role{})
	//config.DB.Create(&entity.Role{
	//	Id:          1,
	//	Name:        "ADMIN",
	//	Description: "管理员角色",
	//})
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
