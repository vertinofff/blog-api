package service_errors

type ServiceError struct {
	Public string
	Err    error
}

func (e *ServiceError) Error() string {
	if e.Public != "" {
		return e.Public
	}
	if e.Err != nil {
		return e.Err.Error()
	}
	return "internal error"
}
func (e *ServiceError) Unwrap() error { return e.Err }
