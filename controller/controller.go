package controller

import (
	"github.com/gin-gonic/gin"
	"MIS_project/pojo"
	"MIS_project/service"
	"strings"
	"fmt"
	"time"
	"strconv"
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
	if service.CheckFix(token) {
		c.JSON(200, map[string]interface{}{
			"msg": "ok",
			"infos": []pojo.InfoBack{{
					ButtonId: "info",
					ObjUrl: "/info",
					ButtonName: "个人信息",
				}, {
					ButtonId: "break",
					ObjUrl: "/break",
					ButtonName: "报修",
				},
			},
		})
		return
	}
	if service.CheckRoot(&token) {
		c.JSON(200, map[string]interface{}{
			"msg": "ok",
			"infos": []pojo.InfoBack{{
					ButtonId: "info",
					ObjUrl: "/info",
					ButtonName: "个人信息",
				}, {
					ButtonId: "room",
					ObjUrl: "/room",
					ButtonName: "寝室管理",
				}, {
					ButtonId: "clean",
					ObjUrl: "/clean",
					ButtonName: "卫生检查登记",
				}, {
					ButtonId: "break",
					ObjUrl: "/break",
					ButtonName: "报修",
				}, {
					ButtonId: "lost",
					ObjUrl: "/lost",
					ButtonName: "失物招领管理",
				}, //TODO
			},
		})
		return
	}
	c.JSON(200, map[string]interface{}{
		"msg": "ok",
		"infos": []pojo.InfoBack{{
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
		},
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

// [#] 失物招领
// [*] get, /lost
// [✓] .
func GetLost(c *gin.Context) {
	token := c.Query("token")
	if !service.CheckToken(&token) {
		c.Redirect(301, "/login")
		return
	}
	if service.CheckRoot(&token) {
		c.HTML(200, "lostroot.html", gin.H{})
		return
	}
	c.HTML(200, "lost.html", gin.H{})

}

// [#] 获取失物招领信息
// [*] post, /lost
// [✓] .
func PostLost(c *gin.Context) {
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
	losts := service.QueryLosts()
	c.JSON(200, map[string]interface{}{
		"msg": "ok",
		"losts": losts,
	})
}

// [#] 改变失物的状态
// [*] post, /lost/change
// [✓] .
func LostChange(c *gin.Context) {
	token, isOk := c.GetPostForm("token")
	if !isOk {
		c.JSON(200, map[string]interface{}{
			"msg": "fail",
		})
		return
	}
	cid, isOk := c.GetPostForm("id")
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
	c.JSON(200, map[string]interface{}{
		"msg": service.ChangeLostionStatus(token, cid),
	})
}

// [#] 新建一个丢失物
// [*] post, /lost/new
// [✓] .
func NewLost(c *gin.Context) {
	// 获取token，验证token
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
	// 获取物品名和图片文件
	name, isOk := c.GetPostForm("name")
	f, err := c.FormFile("file")
	if err != nil {
		c.JSON(200, map[string]interface{}{
			"msg": "fail",
		})
		return
	}
	fmt.Println(name)
	// 将图片保存起来 获取保存之后的url
	ls := strings.Split(f.Filename, ".")
	lst := ls[len(ls) - 1]
	filename := fmt.Sprintf("%d.%s", time.Now().Unix(), lst)
	c.SaveUploadedFile(f, fmt.Sprintf("./static/pic/%s", filename))
	// 将数据存入数据库
	service.AddLostion(name, "/pic/" + filename)
}

// [#] 报修的页面
// [*] get, /break
// [✓] ..
func GetBreak(c *gin.Context) {
	token := c.Query("token")
	if !service.CheckToken(&token) {
		c.Redirect(301, "/login")
		return
	}
	if service.CheckFix(token) {
		c.HTML(200, "breakfix.html", gin.H{})
		return
	}
	c.HTML(200, "break.html", gin.H{})
}

// [#] 返回报修的信息
// [*] post, /break
// [✓] .
func PostBreak(c *gin.Context) {
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
	c.JSON(200, map[string]interface{}{
		"msg": "ok",
		"breaks": service.BreaksInfo(),
	})
}

// [#] 维修报修物品
// [*] post, /break/fix
// [✓] .
func FixBreak(c *gin.Context) {
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
	if !service.CheckFix(token) {
		c.JSON(200, map[string]interface{}{
			"msg": "fail",
		})
		return
	}
	id, isOk := c.GetPostForm("id")
	if !isOk {
		c.JSON(200, map[string]interface{}{
			"msg": "fail",
		})
		return
	}
	c.JSON(200, map[string]interface{}{
		"msg": service.FixBreak(id),
	})
}

// [#] 上报新的报修
// [*] post, /break/new
// [✓] .
func NewBreak(c *gin.Context) {
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
	desc, isOk := c.GetPostForm("desc")
	if !isOk {
		c.JSON(200, map[string]interface{}{
			"msg": "fail",
		})
		return
	}
	fmt.Println(desc)
	service.NewBreak(desc, token)
}

// [#] 返回卫生检查的页面
// [*] get, /clean
// [✓] .
func GetClean(c *gin.Context) {
	token := c.Query("token")
	if !service.CheckToken(&token) {
		c.Redirect(301, "/login")
		return
	}
	if service.CheckRoot(&token) {
		c.HTML(200, "cleanroot.html", gin.H{})
		return
	}
	c.HTML(200, "clean.html", gin.H{})
}

// [#] 返回卫生检查列表
// [*] post, /clean
// [✓] .
func PostClean(c *gin.Context) {
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
	c.JSON(200, map[string]interface{}{
		"msg": "ok",
		"cleans": service.Cleans(),
	})
}

// [#] 查看寝室列表
// [*] post, /room/list
// [✓] ..
func PostRoomList(c *gin.Context) {
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
	c.JSON(200, map[string]interface{}{
		"msg": "ok",
		"rooms": service.RoomList(),
	})
}

// [#] 创建新的卫生检查记录
// [*] post, /clean/new
// [✓] ..
func NewClean(c *gin.Context) {
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
	roomIdS, isOk := c.GetPostForm("room_id")
	if !isOk {
		c.JSON(200, map[string]interface{}{
			"msg": "fail",
		})
		return
	}
	roomIdI, err := strconv.Atoi(roomIdS)
	roomId := uint(roomIdI)
	desc, isOk := c.GetPostForm("desc")
	if !isOk {
		c.JSON(200, map[string]interface{}{
			"msg": "fail",
		})
		return
	}
	scoreS, isOk := c.GetPostForm("score")
	if !isOk {
		c.JSON(200, map[string]interface{}{
			"msg": "fail",
		})
		return
	}
	score, err := strconv.ParseFloat(scoreS, 64)
	if err != nil {
		c.JSON(200, map[string]interface{}{
			"msg": "fail",
		})
		return
	}
	service.NewClean(roomId, desc, score)
	c.JSON(200, map[string]interface{}{
		"msg": "ok",
	})
}
