package error_utils

import "net/http"

type RestErr struct {
	StatusCode       int           `json:"status_code"`
	ErrorCode        int           `json:"error_code"`
	NotificationType string        `json:"notification_type"`
	Message          string        `json:"message"`
	Causes           []interface{} `json:"causes"`
	Error            string        `json:"error"`
}

func NewBadRequestError(message string, errCode int) *RestErr {
	return &RestErr{
		StatusCode:       http.StatusBadRequest,
		ErrorCode:        errCode,
		NotificationType: "warning",
		Message:          message,
		Causes:           nil,
		Error:            "bad_request",
	}
}

func NewUnauthorizedError(message string, errCode int) *RestErr {
	return &RestErr{
		StatusCode:       http.StatusUnauthorized,
		ErrorCode:        errCode,
		NotificationType: "warning",
		Message:          message,
		Causes:           nil,
		Error:            "unauthorized",
	}
}

func NewForbiddenError(message string, errCode int) *RestErr {
	return &RestErr{
		StatusCode:       http.StatusForbidden,
		ErrorCode:        errCode,
		NotificationType: "danger",
		Message:          message,
		Causes:           nil,
		Error:            "forbidden",
	}
}

func NewNotFoundError(message string, errCode int) *RestErr {
	return &RestErr{
		StatusCode:       http.StatusNotFound,
		ErrorCode:        errCode,
		NotificationType: "warning",
		Message:          message,
		Causes:           nil,
		Error:            "not_found",
	}
}

func NewTooManyRequestsError(message string, errCode int) *RestErr {
	return &RestErr{
		StatusCode:       http.StatusTooManyRequests,
		ErrorCode:        errCode,
		NotificationType: "info",
		Message:          message,
		Causes:           nil,
		Error:            "too_many_requests",
	}
}

func NewInternalServerError(message string, errCode int) *RestErr {
	return &RestErr{
		StatusCode:       http.StatusInternalServerError,
		ErrorCode:        errCode,
		NotificationType: "danger",
		Message:          message,
		Causes:           nil,
		Error:            "internal_server_error",
	}
}
