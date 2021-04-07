package request

type PostPageHttpBody struct {
	PageUrl      string  `json:"page_url"`
	UserId       string  `json:"user_id"`
	ApiKey       string  `json:"api_key"`
	ActionType   string  `json:"action_type"`
	HappenTime   int     `json:"happen_time"`
	HappenDay    string  `json:"happen_day"`
	HttpUrl      string  `json:"http_url"`
	LoadTime     float64 `json:"load_time"`
	Status       int     `json:"status"`
	StatusText   string  `json:"status_text"`
	StatusResult string  `json:"status_result"`
	RequestText  string  `json:"request_text"`
	ResponseText string  `json:"response_text"`
	// 设备信息
	IP             string `json:"ip"`
	Device         string `json:"device"`
	DeviceType     string `json:"device_type"`
	Os             string `json:"os"`
	OsVersion      string `json:"os_version"`
	Browser        string `json:"browser"`
	BrowserVersion string `json:"browser_version"`
	UA             string `json:"ua"`
}

// httpinfo get请求参数
//type GetHttpInfoParams struct {
//}
