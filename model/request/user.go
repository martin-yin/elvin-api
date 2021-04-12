package request

type UsersRequest struct {
	SearchDate string `form:"search_date"`
	SearchHour string `form:"search_hour"`
	UserId     string `form:"user_id"`
}
