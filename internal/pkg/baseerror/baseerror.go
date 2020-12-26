package baseerror

import (
	"errors"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

/*
通用错误error
*/

type (
	// BaseError 基本错误类型
	BaseError struct {
		message string
	}

	// ErrorWrapper 定义错误返回
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

// ParamError通用参数验证方法
func ParamError(s interface{}) error {
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")
	validate := validator.New()
	//验证器注册翻译器
	err := zh_translations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		return err
	}
	err = validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return errors.New(err.Translate(trans))
		}
	}
	return nil
}
