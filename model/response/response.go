package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = 7
	SUCCESS = 0
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "操作成功", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}

type StackResult struct {
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
}

type QuotaResult struct {
	Ttfb     float64 `json:"ttfb"`
	DomParse float64 `json:"dom_parse"`
	LoadPage float64 `json:"load_page"`
	Pv       int     `json:"pv"`
}

type WebLoadpageInfo struct {
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

type GetWebLoadPageInfoS struct {
	QuotaResult QuotaResult       `json:"quota_result"`
	StackResult StackResult       `json:"stack_result"`
	PageList    []WebLoadpageInfo `json:"page_list"`
}
