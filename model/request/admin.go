package request

type Login struct {
	Username string `form:"user_name" json:"user_name"`
	Password string `form:"password" json:"password"`
}

type AdminParam struct {
	Username string `form:"user_name" json:"user_name"`
	Password string `form:"password" json:"password"`
	Nickname string `form:"nick_name" json:"nick_name"`
}
