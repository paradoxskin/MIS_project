package pojo

type LoginInfo struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type LoginReturn struct {
	Status string `json:"status"`
	Token string `json:"token"`
}

type Token struct {
	Token string
	Id  uint
}
