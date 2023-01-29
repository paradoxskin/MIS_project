package pojo

type LoginInfo struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type LoginReturn struct {
	Status string `json:"status"`
	Username string `json:"username"`
	Token string `json:"token"`
	Start int64 `json:"start"`
}
