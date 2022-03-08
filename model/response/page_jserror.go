package response

import "time"

type Issues struct {
	ID        string `json:"id"`
	PageUrl   string `json:"page_url"`
	ErrorName string `json:"error_name"`
	Message   string `json:"message"`
	Total     string `json:"total"`
	Today     string `json:"today"`
	ErrorUser string `json:"error_user"`
	LastTime  string `json:"last_time"`
	FirstTime string `json:"first_time"`
}

type PageJsErrDetail struct {
	ID            uint   `json:"id"`
	PageUrl       string `json:"page_url"`
	ComponentName string `json:"componentName"`
	Message       string `json:"message"`
	Stack         string `json:"stack"`
	ErrorName     string `json:"error_name"`
	StackFrames   string `json:"stack_frames"`

	Issues_id       uint      `json:"issues_id"`
	PreviousErrorID uint      `json:"previous_error_id"`
	NextErrorID     uint      `json:"next_error_id"`
	UserId          string    `json:"user_id"`
	MonitorId       string    `json:"monitor_id"`
	ActionType      string    `json:"action_type"`
	HappenTime      int       `json:"happen_time"`
	HappenDay       string    `json:"happen_day"`
	IP              string    `json:"ip"`
	SessionId       string    `json:"session_id"`
	Device          string    `json:"device"`
	DeviceType      string    `json:"device_type"`
	Os              string    `json:"os"`
	OsVersion       string    `json:"os_version"`
	Browser         string    `json:"browser"`
	BrowserVersion  string    `json:"browser_version"`
	UA              string    `json:"ua"`
	Nation          string    `json:"nation"`
	Province        string    `json:"province"`
	City            string    `json:"city"`
	District        string    `json:"district"`
	CreatedAt       time.Time `json:"created_at"`
}
