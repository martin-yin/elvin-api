package request

type StartEndTime struct {
	StartTime string `form:"start_time"`
	EndTime   string `form:"end_time"`
}

type MonitorId struct {
	MonitorId string `form:"monitor_id"`
}
