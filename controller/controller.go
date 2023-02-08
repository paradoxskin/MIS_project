package controller

import (
	"github.com/gin-gonic/gin"
	"MIS_project/pojo"
	"MIS_project/service"
	"fmt"
)

// [#] 返回登陆界面
// [*] get, /login
// [✓] ...
func GetLoginPage(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{})
}

// [#] 用户登陆操作
// [*] post, /login
// [✓] ..
func LoginConfig(c *gin.Context) {
	var login pojo.LoginInfo
	err := c.ShouldBind(&login)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(200, *service.CheckLogin(&login))
}

// [#] 返回登录后的主页界面，区分用户和非用户
// [*] get, /index
// [✓] ..
func GetIndexPage(c *gin.Context) {
	token := c.Query("token")
	if !service.CheckToken(&token) {
		c.Redirect(301, "/login")
		return
	}
	if service.CheckRoot(&token) {
		c.HTML(200, "index.html", gin.H{})
		return
	}
	c.HTML(200, "index.html", gin.H{})// index
}

// [#] 退出，传入token
// [*] post, /quit
// [✓] ...
func QuitWithToken(c *gin.Context) {
	token, isOk := c.GetPostForm("token")
	if isOk {
		service.RemoveToken(&token)
		c.JSON(200, map[string]interface{}{
			"msg": "ok",
		})
		return
	}
	c.JSON(200, map[string]interface{}{
		"msg": "fail",
	})
}
