package service

import (
	"MIS_project/pojo"
	"time"
	"MIS_project/dao"
)

// 返回token
func CheckLogin(loginInfo *pojo.LoginInfo) (*pojo.LoginReturn) {
	var loginReturn pojo.LoginReturn
	// 验证登录
	if loginInfo.Password != dao.QueryPsw(&loginInfo.Username) {
		loginReturn.Status = "fail"
		return &loginReturn
	}
	loginReturn.Status = "success"
	loginReturn.Username = loginInfo.Username
	loginReturn.Token = loginInfo.Password
	loginReturn.Start = time.Now().Unix()
	return &loginReturn
}
