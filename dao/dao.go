package dao

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"MIS_project/pojo"
	"fmt"
)

var DB *gorm.DB

func InitConnect() {
	var err error
	DB, err = gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("[x] database init failed!: %v\n", err))
	}
}

func QueryPsw(username *string) (string, error){
	var user pojo.User
	err := DB.Table("users").Where("user_name = ?", username).Find(&user).Error
	return user.Password, err
}
