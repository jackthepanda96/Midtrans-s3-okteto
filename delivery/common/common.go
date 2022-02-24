package common

import "net/http"

type Structure struct {
	Code    int
	Message string
	Status  bool
	Data    interface{}
}

func StatusGood(code int, message string, data interface{}) Structure {
	return Structure{
		Code:    code,
		Message: message,
		Status:  true,
		Data:    data,
	}
}

func StatusBadRequest(message string) Structure {
	return Structure{
		Code:    http.StatusBadRequest,
		Message: message,
		Status:  false,
		Data:    nil,
	}
}
