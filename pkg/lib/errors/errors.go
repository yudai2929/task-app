package errors

import (
	"database/sql"
	"errors"
	"fmt"
	"runtime/debug"

	"github.com/yudai2929/task-app/database/gen"
	"github.com/yudai2929/task-app/pkg/lib/errors/codes"
)

type customError struct {
	code   codes.Code
	origin error
	stack  string
}

func (c *customError) Error() string {
	if c.origin == nil {
		return fmt.Sprintf("%s: nil", c.code.String())
	}

	return fmt.Sprintf("%s: %s", c.code.String(), c.origin.Error())
}

func New(code codes.Code) error {
	return &customError{
		code:   code,
		origin: fmt.Errorf("%s", code.String()), //nolint:goerr113
		stack:  string(debug.Stack()),
	}
}

func Newf(code codes.Code, format string, a ...interface{}) error {
	return &customError{
		code:   code,
		origin: fmt.Errorf(format, a...), //nolint:goerr113
		stack:  string(debug.Stack()),
	}
}

func Code(err error) codes.Code {
	converted := &customError{} //nolint:exhaustruct
	if !errors.As(err, &converted) {
		return codes.CodeUnknown
	}
	return converted.code
}

func EqualCode(err error, code codes.Code) bool {
	converted := &customError{} //nolint:exhaustruct
	if !errors.As(err, &converted) {
		return false
	}
	return converted.code == code
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}

func Convert(err error) error {
	converted := &customError{} //nolint:exhaustruct
	if errors.As(err, &converted) {
		return Newf(converted.code, "%s", converted.origin.Error())
	}

	fmt.Printf("%+v\n", err)
	if errors.Is(err, gen.ErrAlreadyExists) {
		return Newf(codes.CodeAlreadyExists, "%s", err.Error())
	} else if errors.Is(err, gen.ErrDoesNotExist) {
		return Newf(codes.CodeNotFound, "%s", err.Error())
	} else if errors.Is(err, sql.ErrNoRows) {
		return Newf(codes.CodeNotFound, "%s", err.Error())
	}
	return Newf(codes.CodeUnknown, "%s", err.Error())
}
