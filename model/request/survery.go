package request

type SurveyParams struct {
	TimeGrain string `form:"time_grain"`
	StageType string `form:"stage_type"`

	MonitorId
	StartEndTime
}
