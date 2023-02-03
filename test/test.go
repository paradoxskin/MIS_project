package test

import (
	"MIS_project/pojo"
	"MIS_project/dao"
	"fmt"
)

func Test() {
	dao.DB.AutoMigrate(&pojo.User{})
	var user []pojo.User
	result := dao.DB.Table("users").Where("user_name = ? AND password = ?", "hello", "exxz").Find(&user)
	//dao.DB.Table("users").First(&user)
	fmt.Println("----------")
	fmt.Println(result.RowsAffected)
	fmt.Println("----------")
}
