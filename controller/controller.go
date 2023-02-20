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
	c.HTML(200, "index.html", gin.H{})// index
}

// [#] 根据用户的权限返回应该加载的内容
// [*] post, /index
// [✓] .
func GiveIndexInfo(c *gin.Context) {
	token, isOk := c.GetPostForm("token")
	if !isOk {
		c.JSON(200, map[string]interface{}{
			"msg": "fail",
		})
		return
	}
	if !service.CheckToken(&token) {
		c.JSON(200, map[string]interface{}{
			"msg": "fail",
		})
		return
	}
	if service.CheckRoot(&token) {
		c.JSON(200, map[string]interface{}{
			"msg": "ok",
			// TODO
		})
		return
	}
	infoBacks := []pojo.InfoBack{{
			ButtonId: "info",
			ObjUrl: "/info",
			ButtonName: "个人信息",
		}, {
			ButtonId: "room",
			ObjUrl: "/room",
			ButtonName: "寝室",
		}, {
			ButtonId: "clean",
			ObjUrl: "/clean",
			ButtonName: "卫生检查",
		}, {
			ButtonId: "break",
			ObjUrl: "/break",
			ButtonName: "报修",
		}, {
			ButtonId: "lost",
			ObjUrl: "/lost",
			ButtonName: "失物招领",
		},
	}
	c.JSON(200, map[string]interface{}{
		"msg": "ok",
		"infos": infoBacks,
	})

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

// [#] 信息页面
// [*] get, /info
// [✓] .
func GetInfo(c *gin.Context) {
	token := c.Query("token")
	if !service.CheckToken(&token) {
		c.Redirect(301, "/login")
		return
	}
	c.HTML(200, "info.html", gin.H{})
}

// [#] 信息页面
// [*] post, /info
// [✓] ...
func PostInfo(c *gin.Context) {
	token, isOk := c.GetPostForm("token")
	if !isOk {
		c.JSON(200, map[string]interface{}{
			"msg": "fail",
		})
		return
	}
	if !service.CheckToken(&token) {
		c.JSON(200, map[string]interface{}{
			"msg": "fail",
		})
		return
	}
	// 获取token对应用户的姓名，寝室号，权限
	infos := service.PersonalInfo(token)
	c.JSON(200, map[string]interface{}{
		"msg": "ok",
		"infos": infos,
	})
}
