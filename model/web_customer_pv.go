package model

import "danci-api/global"

type WebCustomerPv struct {
	global.GVA_MODEL

	PageUrl    string `json:"page_url"`
	UserId     string `json:"user_id"`
	UploadType string `json:"upload_type"`
	HappenTime string `json:"happen_time"`
	HappenDate string `json:"happen_date"`
	PageKey    string `json:"page_key"`

	DeviceName     string `json:"device_name"`
	Os             string `json:"os"`
	BrowserName    string `json:"brower_name"`
	BrowserVersion string `json:"brower_version"`
	UA             string `json:"ua"`
	LoadType       string `json:"load_type"`
	LoadTime       string `json:"load_time"`
}
