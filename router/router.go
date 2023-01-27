package router

import (
	"github.com/gin-gonic/gin"
)

func Build() *gin.Engine {
	r := gin.Default()
	return r
}

