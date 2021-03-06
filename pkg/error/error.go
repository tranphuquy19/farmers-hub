package error

import (
	"net/http"
)

type RequestError struct {
	ErrorString string
	ErrorCode   int
}

var (
	ErrInvalidParam  = &RequestError{ErrorString: "Bad Request", ErrorCode: http.StatusBadRequest}
	ErrInternalError = &RequestError{ErrorString: "Internal Error", ErrorCode: http.StatusInternalServerError}
	ErrNotFound      = &RequestError{ErrorString: "Not Found", ErrorCode: http.StatusNotFound}
	ErrUnauthorized  = &RequestError{ErrorString: "Unauthorized", ErrorCode: http.StatusUnauthorized}
	ErrForbidden     = &RequestError{ErrorString: "Forbidden", ErrorCode: http.StatusForbidden}
)

// Code returns the http error code
func (err *RequestError) Code() int {
	return err.ErrorCode
}

func (err *RequestError) Error() string {
	return err.ErrorString
}

func ErrorMessage(errorType *RequestError) (code int, message map[string]interface{}) {
	return errorType.Code(), map[string]interface{}{"error_message": errorType.Error()}
}
