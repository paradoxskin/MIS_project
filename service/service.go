package service

import (
	"MIS_project/pojo"
	"MIS_project/dao"
	"time"
)

// 返回token
func CheckLogin(loginInfo *pojo.LoginInfo) (*pojo.LoginReturn) {
	var loginReturn pojo.LoginReturn
	// 验证登录
	password, err := dao.QueryPsw(&loginInfo.Username)
	if err != nil || loginInfo.Password != password {
		loginReturn.Status = "fail"
		return &loginReturn
	}
	loginReturn.Status = "success"
	loginReturn.Username = loginInfo.Username
	loginReturn.Token = loginInfo.Password
	loginReturn.Start = time.Now().Unix()
	return &loginReturn
}
