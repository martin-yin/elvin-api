package request

type BehaviorRequest struct {
	BehaviorId   string `form:"behavior_id"`
	BehaviorType string `form:"behavior_type"`
}
