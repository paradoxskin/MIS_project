package service

import (
	"MIS_project/pojo"
	"MIS_project/dao"
	"time"
)

// 返回token
func CheckLogin(loginInfo *pojo.LoginInfo) (*pojo.LoginReturn) {
	var loginReturn pojo.LoginReturn
	// 验证登录，没查到为空字符串，不会返回错误信息
	password, err := dao.QueryPsw(&loginInfo.Username)
	if password == "" || err != nil || loginInfo.Password != password {
		loginReturn.Status = "fail"
		return &loginReturn
	}
	loginReturn.Status = "success"
	loginReturn.Username = loginInfo.Username
	loginReturn.Token = loginInfo.Password
	loginReturn.Start = time.Now().Unix()
	return &loginReturn
}
