package service

import (
	"MIS_project/pojo"
	"MIS_project/dao"
	"time"
	"math/rand"
	"fmt"
)

var tokenMap map[string]pojo.Token = make(map[string]pojo.Token)

// 返回token
// [#] 验证用户并返回token
// [*] to controller
// [✓] ..
func CheckLogin(loginInfo *pojo.LoginInfo) (*pojo.LoginReturn) {
	var loginReturn pojo.LoginReturn
	// 验证登录，没查到为空字符串，不会返回错误信息
	password, err := dao.QueryPsw(&loginInfo.Username)
	if password == "" || err != nil || loginInfo.Password != password {
		loginReturn.Status = "fail"
		return &loginReturn
	}
	token := *genRandomToken()
	loginReturn.Status = "good"
	loginReturn.Token = token
	id, err := dao.QueryId(&loginInfo.Username)
	tokenMap[token] = pojo.Token{
		Token: token,
		Id: id,
	}
	return &loginReturn
}

// [#] 随机生成token
// [*] local
// [✓] ...
func genRandomToken() *string {
	var token string
	rand.Seed(time.Now().Unix())
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	for i := 0; i < 18; i++ {
		token += string(chars[rand.Intn(len(chars))])
	}
	return &token
}

// [#] 验证token是否可用
// [*] to controller
// [✓] ...
func CheckToken(token *string) bool {
	_, flag := tokenMap[*token]
	return flag
}

// [#] 验证指定Token对应的用户是否为管理员
// [*] to controller
// [✓] ...
func CheckRoot(token *string) bool {
	id := tokenMap[*token].Id
	power, _ := dao.QueryPower(&id)
	isRoot := power == 1
	return isRoot
}

// [#] 删除token
// [*] to controller
// [✓] ...
func RemoveToken(token *string) {
	delete(tokenMap, *token)
}
// [#] 获取token对应用户的姓名，寝室，权限
// [*] to controller
// [✓] ...
func PersonalInfo(token string) []string {
	id := tokenMap[token].Id
	name, err := dao.QueryName(id)
	if err != nil {
		return []string{}
	}
	isRoot := "普通用户"
	if CheckRoot(&token) {
		isRoot = "管理员"
	} else if CheckFix(token) {
		isRoot = "维修员"
	}
	roomName := dao.QueryRoomName(id)
	return []string {
		name,
		roomName,
		isRoot,
	}
}

// [#] 获取数据库中的失物信息
// [*] to controller
// [✓] ...
func QueryLosts() []pojo.Lostion {
	return dao.QueryLost()
}

// [#] 尝试改变数据库中失物的状态
// [*] to controller
// [✓] ...
func ChangeLostionStatus(token string, cid string) string {
	// 查询物品的状态
	status := dao.QueryLostStatus(cid)
	if status == 2 {
		return "fail"
	}
	if status == 0 {
	// 如果物品状态0，置为1
		dao.ChangeLostStatus(cid, 1)
	} else {
	// 如果物品状态1且用户为root，置为2
		if !CheckRoot(&token) {
			return "fail"
		}
		dao.ChangeLostStatus(cid, 2)
	}
	return "ok"
}

// [#] 向失物数据库中增加物品
// [*] to controller
// [✓] ...
func AddLostion(name string, url string) {
	dao.AddLost(url, name, fmt.Sprintf("%d/%d/%d", time.Now().Year(), time.Now().Month(), time.Now().Day()))
}

// [#] 报修信息
// [*] to controller
// [✓] .
func BreaksInfo() []pojo.Breaks{
	return dao.QueryBreaks()
}

// [#] 检测当前登陆的用户是否为维修员
// [*] to controller
// [✓] ...
func CheckFix(token string) bool {
	id := tokenMap[token].Id
	power, _ := dao.QueryPower(&id)
	isFix := power == 2
	return isFix
}

// [#] 维修指定id的物品
// [*] to controller
// [✓] ...
func FixBreak(id string) string {
	if dao.QueryBreakStatus(id) == 1 {
		return "fail"
	}
	dao.ChangeBreakStatus(id)
	return "ok"
}

func NewBreak(desc string, token string) {
	id := tokenMap[token].Id
	dao.AddBreak(dao.QueryRoomId(id), desc)
}

// [#] 获取卫生检查记录列表
// [*] to controller
// [✓] ...
func Cleans() []pojo.Cleans {
	return dao.QueryClean()
}

// [#] 获取寝室列表
// [*] to controller
// [✓] ...
func RoomList() []pojo.RoomList {
	return dao.QueryRoomList()
}

// [#] 创建新的卫生检查记录并且修改数据库中的得分
// [*] to controller
// [✓] ...
func NewClean(roomId uint, desc string, point float64) {
	dao.AddClean(roomId, desc, point)
	oldScore := dao.QueryRoomScore(roomId)
	newScore := oldScore + point
	dao.UpdateRoomScore(roomId, newScore)
}

// [#] 获取寝室信息
// [*] to controller
// [✓] ..
func RoomInfo(token string) pojo.RoomInfo{
	// 获取当前用户id
	id := tokenMap[token].Id
	// 获取当前用户是哪个寝室的
	roomId := dao.QueryRoomId(id)
	// 获取寝室信息
	return pojo.RoomInfo{Room_name: dao.QueryRoomName(id), Score: dao.QueryRoomScore(roomId), Name: dao.QueryRoomates(roomId)}
}
