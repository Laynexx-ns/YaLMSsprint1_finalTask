package errors

type MyError struct {
	Code    int
	Message string
}

func (err *MyError) Error() string {
	return err.Message
}

func NewCustomError(code int, message string) *MyError {
	return &MyError{Code: code, Message: message}
}

var (
	CalcError = NewCustomError(500, "Internal Calculator Error")
)
