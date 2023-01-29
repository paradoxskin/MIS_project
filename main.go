package main

import (
	"MIS_project/router"
	"MIS_project/dao"
)


func main() {
	dao.InitConnect()
	r := router.Build()
	r.Run(":8080")
}
