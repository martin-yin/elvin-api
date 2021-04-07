package model

import "danci-api/global"

type PageHttp struct {
	global.GVA_MODEL

	PageUrl      string  `json:"page_url"`
	HttpUrl      string  `json:"http_url" gorm:"index:http_url"`
	UserId       string  `json:"user_id"`
	ApiKey       string  `json:"api_key"`
	UploadType   string  `json:"upload_type"`
	HappenTime   string  `json:"happen_time"`
	LoadTime     float64 `json:"load_time"`
	Status       int     `json:"status"`
	StatusText   string  `json:"status_text"`
	StatusResult string  `json:"status_result"`
	RequestText  string  `json:"request_text"`
	ResponseText string  `json:"response_text"  gorm:"type:text"`

	IP             string `json:"ip"`
	Device         string `json:"device"`
	DeviceType     string `json:"device_type"`
	Os             string `json:"os"`
	OsVersion      string `json:"os_version"`
	Browser        string `json:"browser"`
	BrowserVersion string `json:"browser_version"`
	UA             string `json:"ua"`
}

type PageHttpStatistical struct {
	global.GVA_MODEL

	PageUrl      string `json:"page_url"`
	HttpUrl      string `json:"http_url" gorm:"index:http_url"`
	SuccessTotal int    `json:"success_total"`
	FailTotal    int    `json:"fail_total"`
	Total        int    `json:"fail_total"`
}
