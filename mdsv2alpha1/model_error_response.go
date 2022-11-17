package mdsv2alpha1

// ErrorResponse struct for ErrorResponse
type ErrorResponse struct {
	// Optional - http status code
	StatusCode int32 `json:"status_code,omitempty"`
	// Optional - Kafka error code (typically 5 digits)
	ErrorCode int32 `json:"error_code,omitempty"`
	// Optional - Type of error
	Type string `json:"type,omitempty"`
	// Required - Top level error message
	Message string `json:"message"`
	// Optional - List of errors
	Errors []ErrorDetail `json:"errors,omitempty"`
}
