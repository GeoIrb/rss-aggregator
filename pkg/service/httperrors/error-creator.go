package httperrors

// NewError ...
func NewError(status int, err error) error {
	return &httpError{
		Code:    status,
		Message: err.Error(),
	}
}
