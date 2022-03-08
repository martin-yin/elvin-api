package model

type PageView struct {
	MODEL
	PageUrl       string      `json:"page_url"`
	DocumentTitle string      `json:"document_title"`
	Referrer      string      `json:"referrer"`
	Encode        string      `json:"encode"`
	CommonFiles   CommonFiles `json:"common_files"  gorm:"embedded"`
}
