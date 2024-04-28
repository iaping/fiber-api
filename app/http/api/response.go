package api

import (
	"net/http"
)

type Response struct {
	Status int         `json:"-"`
	Code   int         `json:"code"`
	Error  string      `json:"error"`
	Data   interface{} `json:"data"`
}

func NewResponse(data interface{}) Response {
	return Response{
		Status: http.StatusOK,
		Data:   data,
	}
}

func NewErrorResponse(err *Error) Response {
	return Response{
		Status: err.Status,
		Code:   err.Code,
		Error:  err.Text,
	}
}
