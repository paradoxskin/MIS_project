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

// [#] 返回给用户的失物招领信息
// [*] send json
// [✓] ..
type Lostion struct {
	Id uint `json:"id"`
	What string `json:"name"`
	When string `json:"date"`
	Link string `json:"link"`
	Picked string `json:"status"`
}

// [#] 返回给用户的报修信息
// [*] send json
// [✓] ...
type Breaks struct {
	Id uint `json:"id"`
	Room_name string `json:"room"`
	Desc string `json:"desc"`
	Status int `json:"status"`
}
