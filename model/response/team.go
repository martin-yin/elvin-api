package response

type TeamList struct {
	Name     string `json:"name"`
	LeaderId int    `json:"leader_id"`
}
