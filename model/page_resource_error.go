package model

// 页面资源错误
type PageResourceError struct {
	MODEL
	PageUrl     string      `json:"page_url"`
	SourceUrl   string      `json:"source_url"`
	ElementType string      `json:"element_type"`
	Status      string      `json:"status"`
	CommonFiles CommonFiles `json:"common_files"  gorm:"embedded"`
}
