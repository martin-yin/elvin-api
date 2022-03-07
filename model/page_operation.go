package model

// page得点击记录
type PageOperation struct {
	MODEL
	PageUrl     string      `json:"page_url"`
	ClassName   string      `json:"class_name"`
	Placeholder string      `json:"placeholder"`
	Path        string      `json:"path"`
	InputValue  string      `json:"Input_value"`
	TagName     string      `json:"tag_name"`
	InnerText   string      `json:"inner_text" gorm:"type:text"`
	CommonFiles CommonFiles `json:"common_files"  gorm:"embedded"`
}
