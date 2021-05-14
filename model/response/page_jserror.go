package response

type PageJsErrorList struct {
	Message   string `json:"message"`
	Stack     string `json:"stack"`
	ErrorName string `json:"error_name"`
	Users     int    `json:"users"`
}
