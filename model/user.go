package model

// 用户行为记录
type User struct {
	MODEL
	CommonFiles CommonFiles `json:"common_files"  gorm:"embedded"`
}
