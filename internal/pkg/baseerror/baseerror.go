package baseerror

/*
通用错误error
*/

type (
	// BaseError 基本错误类型
	BaseError struct {
		message string
	}

	// ErrorWrapper
	ErrorWrapper struct {
		Error string `json:"errors"`
	}
)

// NewBaseError  初始化基本用户类型
func NewBaseError(message string) *BaseError {
	return &BaseError{message: message}
}

// Error 实现Error
func (e *BaseError) Error() string {

	return e.message
}
