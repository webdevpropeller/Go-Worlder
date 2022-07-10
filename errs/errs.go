// 参照:https://qiita.com/shoichiimamura/items/13199f420ebaf0f0c37c

package errs

import (
	"github.com/pkg/errors"
)

// ErrorType ...
type ErrorType uint

const (
	Unknown ErrorType = iota
	Invalidated
	Unauthorized
	Forbidden
	NotFound
	Conflict
	Failed
)

// ErrorTypeを返すインターフェース
type typeGetter interface {
	Type() ErrorType
}

// ErrorTypeを持つ構造体
type customError struct {
	errorType     ErrorType
	originalError error
}

// New 指定したErrorTypeを持つcustomErrorを返す
func (et ErrorType) New(message string) error {
	return customError{errorType: et, originalError: errors.New(message)}
}

// Errorf 指定したErrorTypeを持つcustomErrorを返す
func (et ErrorType) Errorf(format string, args ...interface{}) error {
	return customError{errorType: et, originalError: errors.Errorf(format, args...)}
}

// Wrap 指定したErrorTypeと与えられたメッセージを持つcustomErrorにWrapする
func (et ErrorType) Wrap(err error, message string) error {
	return customError{errorType: et, originalError: errors.Wrap(err, message)}
}

// Wrapf 指定したErrorTypeと与えられたメッセージを持つcustomErrorにWrapする
func (et ErrorType) Wrapf(err error, format string, args ...interface{}) error {
	return customError{errorType: et, originalError: errors.Wrapf(err, format, args...)}
}

// Error errorインターフェースを実装する
func (e customError) Error() string {
	return e.originalError.Error()
}

// Type typeGetterインターフェースを実装する
func (e customError) Type() ErrorType {
	return e.errorType
}

// Wrap 受け取ったerrorがErrorTypeを持つ場合はそれを引き継いで与えられたエラーメッセージを持つcustomErrorにWrapする
func Wrap(err error, message string) error {
	we := errors.Wrap(err, message)
	if ce, ok := err.(typeGetter); ok {
		return customError{errorType: ce.Type(), originalError: we}
	}
	return customError{errorType: Unknown, originalError: we}
}

// Cause errors.CauseのWrapper
func Cause(err error) error {
	return errors.Cause(err)
}

// GetType ErrorTypeを持つ場合はそれを返し、無ければUnknownを返す
func GetType(err error) ErrorType {
	for {
		if e, ok := err.(typeGetter); ok {
			return e.Type()
		}
		break
	}
	return Unknown
}
