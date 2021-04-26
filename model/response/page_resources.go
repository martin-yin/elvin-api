package response

type PageResourcesResponse struct {
	ResourcesList  []ResourcesListResponse `json:"resources_list"`
	ResourcesQuota ResourcesQuotaResponse  `json:"quota"`
}

type ResourcesListResponse struct {
	PageSourceUrl string `json:"page_source_url"`
	SourceCount   string `json:"source_count"`
	UserCount     string `json:"user_count"`
	PageUrlCount  string `json:"page_url_count"`
	ElementType   string `json:"element_type"`
}

type ResourcesQuotaResponse struct {
	ErrorCount int `json:"error_count"`
	ErrorPage  int `json:"error_page"` // 影响页面次数
	ErrorUser  int `json:"error_user"`
}
