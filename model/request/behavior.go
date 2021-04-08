package request

type UserActionRequest struct {
	ActionID   string `form:"action_id"`
	ActionType string `form:"action_type"`
	StartTime  string `form:"start_time"`
	EndTime    string `form:"end_time"`
}

type UserRequest struct {
	ID string `form:"id"`
}

type UserActionsRequest struct {
	EventID string `form:"event_id"`
}
