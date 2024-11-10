package xerr

import (
	"fmt"
	"github.com/illusory-server/accounts/pkg/errors/codes"

	"github.com/pkg/errors"
)

type Error struct {
	Message error
	Code    codes.Code
	Info    map[string]any
}

func (e *Error) Error() string {
	return e.Message.Error()
}

func New(code codes.Code, msg string, infos ...map[string]any) error {
	info := make(map[string]any)
	for _, inf := range infos {
		for k, v := range inf {
			info[k] = v
		}
	}

	return &Error{
		Message: errors.New(msg),
		Code:    code,
		Info:    info,
	}
}

func Code(err error) codes.Code {
	if err == nil {
		return codes.NotCode
	}
	var e *Error
	if errors.As(err, &e) {
		return e.Code
	}
	return codes.NotCode
}

func Info(err error) map[string]any {
	if err == nil {
		return nil
	}
	var e *Error
	if errors.As(err, &e) {
		return e.Info
	}
	return nil
}

func WithValue(err error, key string, val any) error {
	if err == nil {
		return nil
	}
	var e *Error
	if errors.As(err, &e) {
		if e.Info == nil {
			e.Info = make(map[string]any)
		}
		e.Info[key] = val
	}
	return err
}

func RemoveValue(err error, key string) error {
	if err == nil {
		return nil
	}
	var e *Error
	if errors.As(err, &e) {
		if e.Info != nil {
			delete(e.Info, key)
		}
	}
	return err
}

func GetValue(err error, key string) (any, bool) {
	if err == nil {
		return nil, false
	}
	var e *Error
	if errors.As(err, &e) {
		if e.Info != nil {
			val, ok := e.Info[key]
			return val, ok
		}
	}
	return nil, false
}

func Newf(code codes.Code, format string, a ...interface{}) error {
	return &Error{
		Message: errors.Errorf(format, a...),
		Code:    code,
	}
}

func Wrapf(err error, format string, args ...interface{}) error {
	return Wrap(err, fmt.Sprintf(format, args...))
}

func Cause(err error) error {
	var e *Error
	if errors.As(err, &e) {
		return errors.Cause(e.Message)
	}
	return errors.Cause(err)
}

func Wrap(err error, msg string) error {
	if err == nil {
		return nil
	}

	var e *Error
	if errors.As(err, &e) {
		return &Error{
			Message: errors.Wrap(e.Message, msg),
			Code:    e.Code,
		}
	}

	return &Error{
		Code:    codes.Internal,
		Message: errors.Wrap(err, msg),
	}
}

func Unwrap(err error) error {
	if err == nil {
		return nil
	}

	var e *Error
	if errors.As(err, &e) {
		return &Error{
			Message: errors.Unwrap(e.Message),
			Code:    e.Code,
		}
	}

	return &Error{
		Code:    codes.Internal,
		Message: errors.Unwrap(err),
	}
}

func As(err error, target any) bool {
	var e *Error
	if errors.As(err, &e) {
		return errors.As(e.Message, target)
	}
	return errors.As(err, target)
}

func Is(err error, target error) bool {
	var e *Error
	if errors.As(err, &e) {
		return errors.Is(e.Message, target)
	}
	return errors.Is(err, target)
}
