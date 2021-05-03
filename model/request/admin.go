package request

type Login struct {
	Username string `form:"username"`
	Password string `form:"password"`
}
