package mdsv2alpha1

// ErrorDetail struct for ErrorDetail
type ErrorDetail struct {
	ErrorType string `json:"error_type"`
	Message   string `json:"message,omitempty"`
}
