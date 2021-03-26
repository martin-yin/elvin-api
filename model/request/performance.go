package request

type PagePerformance struct {
	Data WebPerformance
}

type WebPerformance struct {
	Appcache        float64 `json:"appcache"`
	Contentdownload float64 `json:"contentdownload"`
	Dns             float64 `json:"dns"`
	Domparsing      float64 `json:"domparsing"`
	PageUrl         string  `json:"page_url"`
	Redirect        float64 `json:"redirect"`
	Res             float64 `json:"res"`
	Tcp             float64 `json:"tcp" `
	Ttfb            float64 `json:"ttfb"`
}

type PageRequest struct {
	Data WebRequest
}

type WebRequest struct {
	GenerateTime int    `json:"generate_time"`
	Method       string `json:"method" gorm:"comment: 请求方式"`
	HttpType     string `json:"http_type"`
	ElapsedTime  int    `json:"elapsed_time"`
	Code         int    `json:"code" gorm:"comment: 响应状态码"`
	ApiUrl       string `json:"api_url" gorm:"comment: api_url"`
	PageUrl      string `json:"page_url" gorm:"comment: 页面url"`
	Message      string `json:"message"`
	IsError      int    `json:"is_error"`
}

type PageResourcesError struct {
	Data WebResourcesError
}

type WebResourcesError struct {
	PageUrl      string `json:"page_url"`
	ResourceUrl  string `json:"resource_url"`
	GenerateTime int    `json:"generate_time"`
	DomPath      string `json:"dom_path"`
}
