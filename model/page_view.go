package model

type PageView struct {
	MODEL
	PageUrl     string      `json:"page_url"`
	CommonFiles CommonFiles `json:"common_files"  gorm:"embedded"`
}
