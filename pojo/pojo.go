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

// [#] 返回给用户的卫生检查列表
// [*] send json
// [✓] ...
type Cleans struct {
	Room_name string `json:"room_name"`
	Desc string `json:"desc"`
	Points float64 `json:"points"`
}

// [#] 返回的寝室列表
// [*] send json
// [✓] ...
type RoomList struct {
	Room_name string `json:"room_name"`
	ID uint `json:"room_id"`
}

// [#] 寝室信息
// [*] send json
// [✓] ...
type RoomInfo struct {
	Room_name string `json:"room_name"`
	Score float64 `json:"score"`
	Name []string `json:"roomates"`
}

// [#] 寝室信息，不返回成员姓名
// [*] send json
// [✓] ...
type RoomInfo2 struct {
	ID uint `json:"id"`
	Room_name string `json:"room_name"`
	Score float64 `json:"score"`
}

type UserInfo struct {
	UserName string `json:"username"`
	Name string `json:"name"`
	Room_name string `json:"room_name"`
}
