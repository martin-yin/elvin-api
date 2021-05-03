package response

import "danci-api/model"

type LoginResponse struct {
	User      model.Admin `json:"user"`
	Token     string      `json:"token"`
	ExpiresAt int64       `json:"expiresAt"`
}
