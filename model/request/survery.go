package request

type SuveyParams struct {
	TimeGrain string `form:"time_grain"`
	StartTime string `form:"start_time"`
	EndTime   string `form:"end_time"`
	StageType string `form:"stage_type"`
	MonitorId string `form:"monitor_id"`
}
