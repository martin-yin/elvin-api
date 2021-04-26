package response

type PagePerformanceResponse struct {
	QuotaResponse               QuotaResponse                 `json:"quota"`
	StackResponse               StackResponse                 `json:"stack"`
	PagePerformanceListResponse []PagePerformanceListResponse `json:"page_list"`
	StageTimeResponse           []StageTimeResponse           `json:"stage_time"`
	RankingHttListResponse      []RankingHttpListResponse     `json:"ranking_http"`
}

type RankingHttpListResponse struct {
	PageUrl  string `json:"page_url"`
	Total    string `json:"total"`
	LoadPage string `json:"load_page"`
}

type StageTimeResponse struct {
	TimeKey      string  `json:"time_key"`
	Redirect     float64 `json:"redirect"`
	Appcache     float64 `json:"appcache"`
	LookupDomain float64 `json:"lookup_domain"`
	Tcp          float64 `json:"tcp"`
	SslT         float64 `json:"ssl_t"`
	Ttfb         float64 `json:"ttfb"`
	Request      float64 `json:"request"`
	DomParse     float64 `json:"dom_parse"`
	LoadPage     float64 `json:"load_page"`
	LoadEvent    float64 `json:"load_event"`
	Pv           int     `json:"pv"`
}

type StackResponse struct {
	Redirect     float64 `json:"redirect"`
	Appcache     float64 `json:"appcache"`
	LookupDomain float64 `json:"lookup_domain"`
	Tcp          float64 `json:"tcp"`
	SslT         float64 `json:"ssl_t"`
	Ttfb         float64 `json:"ttfb"`
	Request      float64 `json:"request"`
	DomParse     float64 `json:"dom_parse"`
	LoadEvent    float64 `json:"load_event"`
}

type QuotaResponse struct {
	Ttfb     float64 `json:"ttfb"`
	DomParse float64 `json:"dom_parse"`
	LoadPage float64 `json:"load_page"`
	Pv       float64 `json:"pv"`
	Fast     float64 `json:"fast"`
}

type PagePerformanceListResponse struct {
	ID        string  `json:"id"`
	PageUrl   string  `json:"page_url"`
	Request   float64 `json:"request"`
	DomParse  float64 `json:"dom_parse"`
	Ttfb      float64 `json:"ttfb"`
	LoadPage  float64 `json:"load_page"`
	LoadEvent float64 `json:"load_event"`
	LoadType  string  `json:"load_type"`
	Pv        int     `json:"pv"`
}
