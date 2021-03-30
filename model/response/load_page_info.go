package response



type WebLoadPageInfoResponse struct {
	QuotaResponse QuotaResponse       `json:"quota"`
	StackResponse StackResponse       `json:"stack"`
	LoadpageInfoListResponse    []LoadpageInfoListResponse `json:"load_page_info_list"`
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
	Pv       int     `json:"pv"`
	Fast     string  `json:"fast"`
}

type LoadpageInfoListResponse struct {
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

