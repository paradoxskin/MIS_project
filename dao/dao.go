package dao

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"MIS_project/pojo"
	"fmt"
)

var DB *gorm.DB

// [#] 初始化数据库连接池
// [*] init
// [✓] ...
func InitConnect() {
	var err error
	DB, err = gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("[x] database init failed!: %v\n", err))
	}
}

// [#] 根据用户名查找密码
// [*] to service
// [✓] ...
func QueryPsw(username *string) (string, error) {
	var user pojo.User
	err := DB.Table("users").Where("user_name = ?", username).Find(&user).Error
	return user.Password, err
}

// [#] 根据用户名查找ID
// [*] to service
// [✓] ...
func QueryId(username *string) (uint, error) {
	var user pojo.User
	err := DB.Table("users").Where("user_name = ?", *username).Find(&user).Error
	return user.ID, err
}

// [#] 根据id查找用户的权限
// [*] to service
// [✓] ..
func QueryPower(id *uint) (int, error) {
	var user pojo.User
	err := DB.Table("users").Where("id = ?", *id).Find(&user).Error
	return user.Power, err
}
