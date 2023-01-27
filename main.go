package main

import (
	"MIS_project/router"
)


func main() {
	r := router.Build()
	r.Run()
}
