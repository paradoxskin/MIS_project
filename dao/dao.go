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
// [✓] ...
func QueryPower(id *uint) (int, error) {
	var user pojo.User
	err := DB.Table("users").Where("id = ?", *id).Find(&user).Error
	return user.Power, err
}

// [#] 根据id查找用户的姓名
// [*] to service
// [✓] ...
func QueryName(id uint) (string, error) {
	var user pojo.User
	err := DB.Table("users").Where("id = ?", id).Find(&user).Error
	return user.Name, err
}

// [#] 根据id查找用户的用户名
// [*] to service
// [✓] ...
func QueryUsername(id uint) (string, error) {
	var user pojo.User
	err := DB.Table("users").Where("id = ?", id).Find(&user).Error
	return user.UserName, err
}

// [#] 根据id查找用户的寝室名
// [*] to service
// [✓] ...
func QueryRoomName(id uint) string {
	var roomName string
	DB.Raw("select room_name from stu_rooms, rooms where stu_rooms.room_id=rooms.id and user_id=?", id).Scan(&roomName)
	return roomName
}

// [#] 查找数据库中所有的失物信息
// [*] to service
// [✓] ...
func QueryLost() []pojo.Lostion {
	var losts []pojo.Lostion
	DB.Raw("select id,what,\"when\",link,picked from losts").Scan(&losts)
	return losts
}

// [#] 查询数据库中指定失物的状态
// [*] to service
// [✓] ...
func QueryLostStatus(id string) int {
	var lost pojo.Lost
	DB.Table("losts").Where("id = ?", id).Find(&lost)
	return lost.Picked
}

// [#] 改变数据库中失物的状态
// [*] to service
// [✓] ...
func ChangeLostStatus(id string, changeTo int) {
	//DB.Exec("update losts set picked = ? where id = ?", 2, 0)
	DB.Table("losts").Where("id = ?", id).Update("picked", changeTo)
}

// [#] 将物品存入失物数据库中
// [*] to service
// [✓] ...
func AddLost(url string, what string, when string) {
	DB.Create(&pojo.Lost{
		When: when,
		What: what,
		Link: url,
		Picked: 0,
	})
	fmt.Println("2")
}

func QueryBreaks() []pojo.Breaks{
	var breaks []pojo.Breaks
	DB.Raw("select room_name,desc,status from breaks, rooms where rooms.id=breaks.room_id").Scan(&breaks)
	return breaks
}
