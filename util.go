package ccloud

import (
	"fmt"
	"strings"
)

type Errorer interface {
	GetError() *Error
}

// ReplyErr wraps reply and its error and returns the error if they
// are non-nil.
func ReplyErr(resp Errorer, err error) error {
	if err != nil {
		switch err.(type) {
		case *Error:
			return err
		default:
			return WrapCoreErr(err, "reply error")
		}
	}
	// can't return resp.GetError() cuz of https://golang.org/doc/faq#nil_error
	if err := resp.GetError(); err != nil {
		return err
	}
	return nil
}

// formatURL resolves a url from the provided format string and provided arguments
func formatURL(endpoint string, arguments ...interface{}) string {
	return fmt.Sprintf(string(endpoint), arguments...)
}

// isValidResource checks for invalid resources that are guaranteed not to be found in the backend.
// If a resource is invalid, we don't need to send a request (see CLI-1637).
func isValidResource(resource string) bool {
	return resource != "" && !strings.HasSuffix(resource, "\r")
}
