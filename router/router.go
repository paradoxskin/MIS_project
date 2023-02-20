package router

import (
	"github.com/gin-gonic/gin"
	"MIS_project/controller"
)

// [#] 初始化
// [*] init
// [✓] ..
func Build() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templ/*")
	r.Static("/js", "static/js")
	r.Static("/css", "static/css")
	login := r.Group("/login")
	{
		login.GET("", controller.GetLoginPage)
		login.POST("", controller.LoginConfig)
	}
	index := r.Group("/index")
	{
		index.GET("", controller.GetIndexPage)
		index.POST("", controller.GiveIndexInfo)
	}
	info := r.Group("/info")
	{
		info.GET("", controller.GetInfo)
		info.POST("", controller.PostInfo)
	}
	lost := r.Group("/lost")
	{
		lost.GET("", controller.GetLost)
	}
	quit := r.Group("/quit")
	{
		quit.POST("", controller.QuitWithToken)
	}
	return r
}

