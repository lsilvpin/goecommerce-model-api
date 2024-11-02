package domain

type ReturnModel struct {
	Trace             string `json:"trace"`
	Message           string `json:"message"`
	HttpStatusMessage string `json:"http_status_message"`
	HttpStatusCode    int    `json:"http_status_code"`
}
