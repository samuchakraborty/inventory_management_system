package utils

// HTTPError example
type HTTPError struct {
	Code    int `json:"code" example:"400"`
	Message any `json:"message" example:"status bad request"`
}
