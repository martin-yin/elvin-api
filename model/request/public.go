package request

type MonitorIdAndTime struct {
	StartTime string `form:"start_time"`
	EndTime   string `form:"end_time"`
	MonitorId string `form:"monitor_id"`
}
