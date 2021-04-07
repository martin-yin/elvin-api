package request

type PostBehaviorInfoBody struct {
	PageUrl     string `json:"page_url"`
	UserId      string `json:"user_id"`
	ApiKey      string `json:"api_key"`
	UploadType  string `json:"upload_type"`
	HappenTime  string `json:"happen_time"`
	ActionType  string `json:"action_type"`
	ClassName   string `json:"class_name"`
	Placeholder string `json:"placeholder"`
	InputValue  string `json:"Input_value"`
	TagNameint  string `json:"tag_name"`
	InnterText  string `json:"innter_text"`
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
