package ccloud

import (
	"fmt"
	"net/http"
	"strings"

	corev1 "github.com/confluentinc/cc-structs/kafka/core/v1"
)

var ErrNotFound = &corev1.Error{Code: http.StatusNotFound, Message: "resource not found"}

// WrapErr returns a standard typed error if one exists, or a generic wrapped corev1.Error with the msg
func WrapErr(err error, msg string) error {
	if err == nil {
		return nil
	}
	switch err.Error() {
	// Maps to these error messages:
	// https://github.com/confluentinc/cc-auth-service/blob/06db0bebb13fb64c9bc3c6e2cf0b67709b966632/jwt/token.go#L23
	// TODO: add app error codes to our API to make this less fragile (more granular than HTTP codes)
	case "token is expired":
		return &ExpiredTokenError{}
	case "malformed token":
		return &InvalidTokenError{Message: err.Error()}
	}
	if strings.Contains(err.Error(), "Token parsing error") {
		return &InvalidTokenError{Message: err.Error()}
	}
	return corev1.WrapErr(err, msg)
}

type InvalidLoginError struct{}

func (e *InvalidLoginError) Error() string {
	return "username or password is invalid"
}

type SuspendedOrganizationError struct{}

func (e *SuspendedOrganizationError) Error() string {
	return "organization is suspended"
}

type UnknownLoginError struct{}

func (e *UnknownLoginError) Error() string {
	return "encountered unexpected error logging in; please try again"
}

type ExpiredTokenError struct{}

func (e *ExpiredTokenError) Error() string {
	return "token is expired"
}

type InvalidTokenError struct {
	Message string
}

func (e *InvalidTokenError) Error() string {
	return e.Message
}

type UnknownAPIKeyError struct {
	APIKey string
}

func (e *UnknownAPIKeyError) Error() string {
	return fmt.Sprintf("Unknown API key %s", e.APIKey)
}

type UnknownConnectorIdError struct {
	ConnectorId string
	Found       string
}

func (e *UnknownConnectorIdError) Error() string {
	return fmt.Sprintf("Unknown Connector ID:  %s. Found: %s", e.ConnectorId, e.Found)
}

func IsUnknownAPIKeyError(err error) bool {
	_, ok := err.(*UnknownAPIKeyError)
	return ok
}
