package ccloud

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
			return WrapErr(err, "reply error")
		}
	}
	// can't return resp.GetError() cuz of https://golang.org/doc/faq#nil_error
	if err := resp.GetError(); err != nil {
		return err
	}
	return nil
}
