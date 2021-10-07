package model

type Admin struct {
	MODEL
	UserName string `json:"user_name" gorm:"unique"`
	Password string `json:"password" sql:"-"`
	Email    string `json:"email"`
	NickName string `json:"nick_name" gorm:"unique"`
}
