package request

type PostPageResourceErroBody struct {
	PageUrl     string `json:"page_url"`
	UserId      string `json:"user_id"`
	ApiKey      string `json:"api_key"`
	UploadType  string `json:"upload_type"`
	HappenTime  string `json:"happen_time"`
	SourceUrl   string `json:"source_url"`
	ElementType string `json:"element_type"`
	Status      string `json:"status"`
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
