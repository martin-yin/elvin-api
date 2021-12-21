package model

// 用户行为记录
type IPAddress struct {
	MODEL
	IP       string `json:"ip"`
	ISP      string `json:"isp"`
	CityId   int64 `json:"city_id"`
	Country  string `json:"country"`
	Region   string `json:"region"`
	Province string `json:"province"`
	City     string `json:"city"`
}
