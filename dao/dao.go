package dao

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"fmt"
)

var db *gorm.DB

func InitConnect() {
	var err error
	db, err = gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("[x] database init failed!: %v\n", err))
	}
}

func QueryPsw(username *string) string {
	
	return ""
}
