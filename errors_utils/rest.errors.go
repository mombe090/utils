package errors_utils

import "net/http"

type RestError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

func GetBadRequest(msg string) *RestError {
	return &RestError{
		Status:  http.StatusBadRequest,
		Message: msg,
		Error:   "BAD_REQUEST",
	}
}

func GetInternatServerError(msg string) *RestError {
	return &RestError{
		Status:  http.StatusInternalServerError,
		Message: msg,
		Error:   "INTERNAL_SERVER_ERROR",
	}
}
