package service

import (
	"MIS_project/pojo"
	"MIS_project/dao"
	"time"
	"math/rand"
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
// [#] 获取token对应用户的用户名，姓名，寝室，权限
// [*] to controller
// [✓] .
func PersonalInfo(token string) []string {
	id := tokenMap[token].Id
	username, err := dao.QueryUsername(id)
	if err != nil {
		return []string{}
	}
	name, err := dao.QueryName(id)
	if err != nil {
		return []string{}
	}
	isRoot := "普通用户"
	if CheckRoot(&token) {
		isRoot = "管理员"
	}
	roomName := dao.QueryRoomName(id)
	return []string {
		username,
		name,
		roomName,
		isRoot,
	}
}
