package codes

import (
	"net/http"
)

type Code int

const (
	CodeUnknown Code = iota
	CodeAborted
	CodeAlreadyExists
	CodeAlreadyExistsSafety
	CodeCancelled
	CodeFailedPrecondition
	CodeInternal
	CodeInvalidArgument
	CodeJWTExpired
	CodeNotFound
	CodeNotFoundSafety
	CodePermissionDenied
	CodeUnauthenticated
	CodeUnimplemented
)

func (c Code) String() string {
	switch c {
	case CodeAborted:
		return "aborted"
	case CodeAlreadyExists, CodeAlreadyExistsSafety:
		return "already exists"
	case CodeCancelled:
		return "cancelled"
	case CodeFailedPrecondition:
		return "failed precondition"
	case CodeInternal:
		return "internal"
	case CodeInvalidArgument:
		return "invalid argument"
	case CodeJWTExpired:
		return "jwt expired"
	case CodeNotFound, CodeNotFoundSafety:
		return "not found"
	case CodePermissionDenied:
		return "permission denied"
	case CodeUnauthenticated:
		return "unauthenticated"
	case CodeUnimplemented:
		return "unimplemented"
	case CodeUnknown:
		return "unknown"
	default:
		return "unknown"
	}
}

func (c Code) HTTPStatus() int {
	switch c {
	case CodeAborted:
		return http.StatusConflict
	case CodeAlreadyExists, CodeAlreadyExistsSafety:
		return http.StatusConflict
	case CodeCancelled:
		return http.StatusRequestTimeout
	case CodeFailedPrecondition:
		return http.StatusBadRequest
	case CodeInternal:
		return http.StatusInternalServerError
	case CodeInvalidArgument:
		return http.StatusBadRequest
	case CodeJWTExpired:
		return http.StatusUnauthorized
	case CodeNotFound, CodeNotFoundSafety:
		return http.StatusNotFound
	case CodePermissionDenied:
		return http.StatusForbidden
	case CodeUnauthenticated:
		return http.StatusUnauthorized
	case CodeUnimplemented:
		return http.StatusNotImplemented
	case CodeUnknown:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}
