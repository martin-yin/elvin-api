package response

type PageJsErrorList struct {
	Message string `json:"message"`
	Stack string `json:"stack"`
	Frequency int `json:"frequency"`
	ID        int `json:"id"`
}
