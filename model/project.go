package model

import "danci-api/global"

type Project struct {
	global.GVA_MODEL

	ProjectName string `json:"project"`
	AppKey string `json:"app_key"`

	AdminId string `json:"admin_id"`
	// bug多的时候通知得邮箱
	Email string `json:"email"`
	//webMonitorId	varchar	36	0	-1	0	0	0							utf8	utf8_unicode_ci	0	0	0	0	0	0	0
	//projectType	varchar	30	0	-1	0	0	0							utf8	utf8_unicode_ci	0	0	0	0	0	0	0
	//projectName	varchar	20	0	-1	0	0	0							utf8	utf8_unicode_ci	0	0	0	0	0	0	0
	//userId	varchar	200	0	-1	0	0	0							utf8	utf8_unicode_ci	0	0	0	0	0	0	0
	//userTag	text	0	0	-1	0	0	0							utf8	utf8_unicode_ci	0	0	0	0	0	0	0
	//email	varchar	200	0	-1	0	0	0							utf8	utf8_unicode_ci	0	0	0	0	0	0	0
	//startList	varchar	100	0	-1	0	0	0							utf8	utf8_unicode_ci	0	0	0	0	0	0	0
	//monitorCode	text	0	0	-1	0	0	0							utf8	utf8_unicode_ci	0	0	0	0	0	0	0
	//fetchCode	text	0	0	-1	0	0	0							utf8	utf8_unicode_ci	0	0	0	0	0	0	0
	//filterDomain	varchar	200	0	-1	0	0	0							utf8	utf8_unicode_ci	0	0	0	0	0	0	0
	//filterType	varchar	20	0	-1	0	0	0							utf8	utf8_unicode_ci	0	0	0	0	0	0	0
	//recording	varchar	2	0	-1	0	0	0							utf8	utf8_unicode_ci	0	0	0	0	0	0	0
	//createdAt	datetime	0	0	-1	0	0	0									0	0	0	0	0	0	0
	//updatedAt	datetime	0	0	-1	0	0	0									0	0	0	0	0	0	0

}
