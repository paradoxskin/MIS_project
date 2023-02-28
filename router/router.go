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
	r.Static("/pic", "static/pic")
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
		lost.POST("", controller.PostLost)
		lost.POST("/change", controller.LostChange)
		lost.POST("/new", controller.NewLost)
	}
	bk := r.Group("/break")
	{
		bk.GET("", controller.GetBreak)
		bk.POST("", controller.PostBreak)
		bk.POST("/fix", controller.FixBreak)
		bk.POST("/new", controller.NewBreak)
	}
	clean := r.Group("/clean")
	{
		clean.GET("", controller.GetClean)
		clean.POST("", controller.PostClean)
		clean.POST("/new", controller.NewClean)
	}
	room := r.Group("/room")
	{
		room.POST("/list", controller.PostRoomList)
		room.GET("", controller.GetRoom)
		room.POST("", controller.PostRoom)
		room.POST("/new", controller.NewRoom)
		room.POST("/mates", controller.Roomates)
	}
	user := r.Group("/user")
	{
		user.GET("", controller.GetUser)
		user.POST("", controller.PostUser)
		user.POST("/new", controller.NewUser)
	}
	quit := r.Group("/quit")
	{
		quit.POST("", controller.QuitWithToken)
	}
	return r
}

