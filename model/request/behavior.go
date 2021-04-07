package request

type UserActionRequest struct {
	ActionID   string `form:"action_id"`
	ActionType string `form:"action_type"`
}
