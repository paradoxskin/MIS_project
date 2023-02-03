package main

import (
	"MIS_project/router"
	"MIS_project/dao"
	"MIS_project/test"
)


func main() {
	dao.InitConnect()
	test.Test()
	r := router.Build()
	r.Run(":8080")
}
