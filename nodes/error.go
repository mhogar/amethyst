package nodes

const (
	ERROR_TYPE_CLIENT   = iota
	ERROR_TYPE_INTERNAL = iota
)

type Error struct {
	Type   int
	Errors []error
}

func ClientError(errs ...error) *Error {
	return &Error{
		Type:   ERROR_TYPE_CLIENT,
		Errors: errs,
	}
}

func InternalError(errs ...error) *Error {
	return &Error{
		Type:   ERROR_TYPE_INTERNAL,
		Errors: errs,
	}
}
