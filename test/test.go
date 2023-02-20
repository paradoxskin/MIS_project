package test

import (
	"MIS_project/pojo"
	"MIS_project/dao"
	"fmt"
)

func Test() {
	//dao.DB.Table("users").First(&user)
	fmt.Println("----------")
	dao.DB.AutoMigrate(&pojo.User{})
	dao.DB.AutoMigrate(&pojo.Room{})
	dao.DB.AutoMigrate(&pojo.StuRoom{})
	var tmp string
	dao.DB.Raw("select room_name from stu_rooms, rooms where stu_rooms.room_id=rooms.id and user_id=?", "1").Scan(&tmp)
	fmt.Println(tmp)
	//dao.DB.Create(&pojo.StuRoom{UserId: 1, RoomId: 1})
	//dao.DB.Create(&pojo.User{UserName: "root", Password: "toor", Name: "ROOT", Power: 1})
	//dao.DB.Create(&pojo.Room{Score: 1, Room_name: "F3南311"})
	fmt.Println("----------")
}
