package web

import (
	"strconv"
)

type Response struct {
	Code string `json:"code"`
	Data interface{} `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

func NewResponse(code int, data interface{}, err string) Response {
	if code < 300{
		return Response{
            Code: strconv.Itoa(code),
            Data: data,
            Error: "",
		}
	}
	return Response{
		Code: strconv.Itoa(code),
            Data: nil,
            Error: err,
	}
}