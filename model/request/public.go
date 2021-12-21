package request

type StartEndTime struct {
	StartTime string `form:"start_time"`
	EndTime   string `form:"end_time"`
}

type MonitorId struct {
	MonitorId string `form:"monitor_id"`
}

type RequestParams struct {
	StartTime string `form:"start_time"`
	EndTime   string `form:"end_time"`
	MonitorId string `form:"monitor_id"`
}
