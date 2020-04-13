package utils

type ApplicationError struct {
	Message   string `json:"message"`
	ErrorCode int  `json:"errorInfo"`
	Code      string `json:"code"`
}
