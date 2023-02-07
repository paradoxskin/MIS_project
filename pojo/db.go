package pojo

import (
	"gorm.io/gorm"
)

// [#] 用户数据表
// [*] database
// [✓] ...
type User struct {
	gorm.Model
	UserName string
	Password string
	Name string
	Power int
}

// [#] 寝室数据表
// [*] database
// [✓] ...
type Room struct {
	gorm.Model
	Room_name string
	Score float64
}

// [#] 卫生检查表
// [*] database
// [✓] ...
type Clean struct {
	gorm.Model
	RoomId uint
	// 时间戳
	When int64 
	Desc string
	Points float64
}

// [#] 保修
// [*] database
// [✓] ...
type Break struct {
	gorm.Model
	PosterId uint
	Desc string
	// 0表示上报，1表示处理完成
	Status int
	PostTime int64
}

// [#] 访客表
// [*] database
// [✓] ...
type Visit struct {
	gorm.Model
	Who string
	// 0表示未离开，1表示离开
	IsLeave int
	WhenArrive int64
	WhenLeave int64
	Why string
}

// [#] 失物招领表
// [*] databse
// [✓] ...
type Lost struct {
	gorm.Model
	When int64
	What string
	Link string
	// 0表示未被取走，1表示已被取走
	Picked int
	PickerId uint
}

// [#] 学生寝室表
// [*] database
// [✓] ...
type StuRoom struct {
	gorm.Model
	UserId int64
	RoomId int64
}
