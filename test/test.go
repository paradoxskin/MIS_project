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
	dao.DB.AutoMigrate(&pojo.Lost{})
	dao.DB.AutoMigrate(&pojo.Break{})
	dao.DB.AutoMigrate(&pojo.Clean{})
	var tmp string
	dao.DB.Raw("select room_name from stu_rooms, rooms where stu_rooms.room_id=rooms.id and user_id=?", "1").Scan(&tmp)
	fmt.Println(tmp)
	fmt.Println(dao.QueryBreaks())
	//fmt.Println(dao.QueryRoomScore(1))
	//dao.UpdateRoomScore(1, 20)
	//fmt.Println(dao.QueryRoomScore(1))
	//service.FixBreak("1")
	//dao.DB.Create(&pojo.User{UserName: "fixer", Password: "fixer", Name: "维修员", Power: 2})
	//dao.DB.Create(&pojo.Clean{Desc: "卫生很整洁", Points: 3, RoomId: 1})
	//dao.DB.Create(&pojo.Break{RoomId: 1, Desc: "花洒坏了", Status: 0})
	//dao.ChangeLostStatus("1", 2)
	//dao.DB.Create(&pojo.Lost{When: "2023/2/19", What: "模型", Link:	"https://i.redd.it/0wll22m17tm41.jpg", Picked: 2})
	//dao.DB.Create(&pojo.Lost{When: "2023/2/21", What: "水杯", Link:	"https://img.zcool.cn/community/011d3c5a58e473a80120121f1940ea.jpg@1280w_1l_2o_100sh.jpg", Picked: 0})
	//dao.DB.Create(&pojo.StuRoom{UserId: 1, RoomId: 1})
	//dao.DB.Create(&pojo.User{UserName: "root", Password: "toor", Name: "ROOT", Power: 1})
	//dao.DB.Create(&pojo.Room{Score: 1, Room_name: "F3南311"})
	fmt.Println("----------")
}
