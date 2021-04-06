package response

type PageJsErrorList struct {
	Message   string `json:"message"`
	Frequency int    `json:"frequency"`
	ID        int    `json:"id"`
}
