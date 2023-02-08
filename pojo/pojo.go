package pojo

// [#] 从form表单获取的用户登录信息
// [*] form
// [✓] ...
type LoginInfo struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

// [#] 返回给用户的登录结果
// [*] sent json
// [✓] ...
type LoginReturn struct {
	Status string `json:"status"`
	Token string `json:"token"`
}

// [#] 返回给用户的主页显示信息
// [*] send json
// [✓] ..
type InfoBack struct {
	ButtonId string `json:"button_id"`
	ObjUrl string `json:"obj_url"`
	ButtonName string `json:"button_name"`
}

type Token struct {
	Token string
	Id  uint
}
