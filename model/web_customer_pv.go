package model

import "danci-api/global"

type WebCustomerPv struct {
	global.GVA_MODEL

	DeviceName     string `json:"device_name"`
	Os             string `json:"os"`
	BrowserName    string `json:"brower_name"`
	BrowserVersion string `json:"brower_version"`
	UA             string `json:"ua"`
	LoadType       string `json:"load_type"`
	LoadTime       string `json:"load_time"`
}
