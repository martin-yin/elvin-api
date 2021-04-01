package response

type WebResourcesInfoResponse struct {
	ResourcesInfoList []ResourcesInfoListResponse `json:"resources_info_list"`
	ResourcesQuota    ResourcesQuota              `json:"resources_quota"`
}

type ResourcesInfoListResponse struct {
	PageSourceUrl string `json:"page_source_url"`
	SourceCount   string `json:"source_count"`
	UserCount     string `json:"user_count"`
	PageUrlCount  string `json:"page_url_count"`
	ElementType   string `json:"element_type"`
}

type ResourcesQuota struct {
	ErrorCount int `json:"error_count"`
	ErrorPage  int `json:"error_page"` // 影响页面次数
	ErrorUser  int `json:"error_user"`
}
