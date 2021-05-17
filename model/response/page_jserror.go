package response

type PageJsErrorList struct {
	ID         string `json:"id"`
	ErrorName  string `json:"error_name"`
	Message    string `json:"message"`
	ErrorCount string `json:"error_count"`
	ErrorUser  string `json:"error_user"`
}

type PageJsErrorDetail struct {
	ID            uint   `json:"id"`
	PageUrl       string `json:"page_url"`
	ComponentName string `json:"componentName"`
	Message       string `json:"message"`
	Stack         string `json:"stack"`
	ErrorName     string `json:"error_name"`
	//PublicFiles   PublicFiles `json:"public_files" gorm:"embedded"`
	JsIssuesId      uint   `json:"js_issues_id"`
	PreviousErrorID uint   `json:"previous_error_id"`
	NextErrorID     uint   `json:"next_error_id"`
	UserId          string `json:"user_id"`
	MonitorId       string `json:"monitor_id"`
	ActionType      string `json:"action_type"`
	HappenTime      int    `json:"happen_time"`
	HappenDay       string `json:"happen_day"`
	IP              string `json:"ip"`
	EventId         string `json:"event_id"`
	Device          string `json:"device"`
	DeviceType      string `json:"device_type"`
	Os              string `json:"os"`
	OsVersion       string `json:"os_version"`
	Browser         string `json:"browser"`
	BrowserVersion  string `json:"browser_version"`
	UA              string `json:"ua"`
	Nation          string `json:"nation"`
	Province        string `json:"province"`
	City            string `json:"city"`
	District        string `json:"district"`
	//componentName: "component <n>"
	//created_at: "2021-05-17T17:20:50+08:00"
	//error_name: "TypeError"
	//id: 21
	//js_issues_id: 1
	//message: "Cannot read property 'length' of null(v-on handler)"
	//page_url: "http://192.168.0.109:8020/#/"
	//stack: "TypeError: Cannot read property 'length' of null\n    at a.value (http://192.168.0.109:8020/js/app.b38e3c6f.js:1:7653)\n    at He (https://cdn.bootcss.com/vue/2.6.11/vue.min.js:6:11384)\n    at HTMLImageElement.n (https://cdn.bootcss.com/vue/2.6.11/vue.min.js:6:13168)\n    at HTMLImageElement.Yr.o._wrapper (https://cdn.bootcss.com/vue/2.6.11/vue.min.js:6:48505)"
	//updated_at: "2021-05-17T17:20:50+08:00"

}
