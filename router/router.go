package router

import (
	"github.com/gin-gonic/gin"
	"MIS_project/controller"
)

func Build() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templ/*")
	login := r.Group("/login")
	{
		login.GET("", controller.GetLoginPage)
		login.POST("", controller.LoginConfig)
	}
	return r
}

