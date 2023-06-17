package helper

type ApiResponse struct {
	Status       string `json:"status"`
	ErrorMessage string `json:"errorMessage"`
	Items        any    `json:"items"`
}
