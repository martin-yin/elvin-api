package services

import (
	"danci-api/global"
	"danci-api/model"
	"danci-api/model/response"
	"fmt"
)

func SetWebLoadPageInfo(weLoadPageInfo model.WebLoadpageInfo) {

	err := global.GVA_DB.Create(&weLoadPageInfo).Error
	if err != nil {
		fmt.Print(err, "err!!!!!!!!!!!!! \n")
	}
}

func WebHttpInfoModel(weLoadPageInfo []*model.WebHttpInfo) {
	err := global.GVA_DB.Create(&weLoadPageInfo).Error
	if err != nil {
		fmt.Print(err, "err!!!!!!!!!!!!! \n")
	}
}

func SetWebResourcesError(webResourceErrorInfo model.WebResourceErrorInfo) {
	err := global.GVA_DB.Create(&webResourceErrorInfo).Error
	if err != nil {
		fmt.Print(err, "err !!!!!!!! \n")
	}
}

func SetBehaviorInfo(webBehaviorInfo model.WebBehaviorInfo) {
	err := global.GVA_DB.Create(&webBehaviorInfo).Error
	if err != nil {
		fmt.Print(err, "err !!!!!!!! \n")
	}
}

func GetWebLoadPageInfo() *response.GetWebLoadPageInfoS {
	var stacking response.StackResult
	var quota response.QuotaResult
	var pageInfoList []response.WebLoadpageInfo
	db := global.GVA_DB.Model(&model.WebLoadpageInfo{})
	db.Select("ID, page_url, request, dom_parse, ttfb, load_page, load_event, load_type, load_page, COUNT(page_url) as pv").Group("page_url").Scan(&pageInfoList)
	db.Select("round(AVG(redirect),2) as redirect, round(AVG(appcache),2) as appcache, round(AVG(lookup_domain),2) as lookup_domain,  round(AVG(tcp),2) as tcp, round(AVG(ssl_t),2) as ssl_t,  round(AVG(request),2) as request, round(AVG(ttfb),2) as ttfb, round(AVG(load_event),2) as load_event").Scan(&stacking)
	db.Select("round(AVG(dom_parse),2) as dom_parse, round(AVG(ttfb),2) as ttfb, round(AVG(load_page),2) as load_page, Count(id) as Pv").Scan(&quota)
	return &response.GetWebLoadPageInfoS{
		QuotaResult: quota,
		StackResult: stacking,
		PageList:    pageInfoList,
	}
}
