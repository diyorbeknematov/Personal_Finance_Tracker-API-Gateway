package models

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}
