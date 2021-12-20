package request

type UsersRequest struct {
	StartTime string `form:"start_time"`
	EndTime   string `form:"end_time"`
	SessionId string `form:"session_id"`
	UserId    string `form:"user_id"`
	MonitorId string `form:"monitor_id"`
}
