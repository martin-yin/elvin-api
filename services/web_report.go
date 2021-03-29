package services

import (
	"danci-api/global"
	"danci-api/model"
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

//appCacheTime: '0',
//dnsTime: '0',
//redirectTime: '0',
//reqTime: '0',
//resTime: '0',
//tcpTime: '0',
//transferSize: '0',
//domParse: '0'
type Result struct {
	Redirect     float64
	Appcache     float64
	LookupDomain float64
	Tcp float64
	Request float64
	DomParse float64
}

func GetWebLoadPageInfo() Result {
	var result Result
	db := global.GVA_DB.Model(&model.WebLoadpageInfo{})
	db.Select("AVG(redirect) as redirect, AVG(appcache) as appcache, AVG(lookup_domain) as lookup_domain, AVG(tcp) as tcp, AVG(request) as request,  AVG(dom_parse) as dom_parse").Scan(&result)
	return result
}
