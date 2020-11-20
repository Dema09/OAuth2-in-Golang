package response_body

import "net/http"

type Response struct{
	StatusCode int `json: "status_code"`
	Message string `json: "message"`
	Data interface{} `json: "data"`
}

func NewStatusOK(data interface{}) *Response {
	return &Response{
		StatusCode: http.StatusOK,
		Message: "OK",
		Data: data,
	}
}

func NewStatusInternalServerError(message string) *Response {
	return &Response{
		StatusCode: http.StatusInternalServerError,
		Message: message,
		Data: nil,
	}
}

func NewStatusBadRequest(message string) *Response{
	return &Response{
		StatusCode: http.StatusBadRequest,
		Message: message,
		Data: nil,
	}
}

func NewStatusNotFound(message string) *Response{
	return &Response{
		StatusCode: http.StatusNotFound,
		Message: message,
		Data: nil,
	}
}

func NewStatusUnauthorized(message string) *Response{
	return &Response{
		StatusCode: http.StatusUnauthorized,
		Message: message,
		Data: nil,
	}
}
