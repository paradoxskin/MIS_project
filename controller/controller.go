package controller

import (
	"github.com/gin-gonic/gin"
	"MIS_project/pojo"
	"MIS_project/service"
	"fmt"
)

func GetLoginPage(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{})
}

func LoginConfig(c *gin.Context) {
	var login pojo.LoginInfo
	err := c.ShouldBind(&login)
	if err != nil {
		// 报错
		fmt.Println(err)
		return
	}
	c.JSON(200, *service.CheckLogin(&login))
}
