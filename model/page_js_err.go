package model

type PageJsErr struct {
	MODEL
	PageUrl       string      `json:"page_url"`
	ComponentName string      `json:"componentName"`
	Message       string      `json:"message"`
	Stack         string      `json:"stack" gorm:"type:text"`
	StackFrames   string      `json:"stack_frames" gorm:"type:text"`
	ErrorName     string      `json:"error_name"`
	CommonFiles   CommonFiles `json:"common_files"  gorm:"embedded"`
	IssuesId      uint        `json:"issues_id"`
}

type Issue struct {
	MODEL
	ErrorName  string      `json:"error_name"`
	Message    string      `json:"message"`
	Stack      string      `json:"stack" gorm:"type:text" gorm:"uniqueIndex"`
	HappenTime int         `json:"happen_time"`
	MonitorId  string      `json:"monitor_id"`
	IsFix      bool        `json:"is_fix" gorm:"type: bool" gorm:"default: false"`
	FixTime    int         `json:"fix_time"`
	FixUserId  int         `json:"fix_user_id"`
	PageJsErr  []PageJsErr `json:"page_js_err" gorm:"foreignKey:IssuesId"`
}
