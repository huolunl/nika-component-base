/*
   @Author:huolun
   @Date:2022/3/15
   @Description
*/
package http_response_base

import "net/http"

var OKResponseBase = NewResponseBase(http.StatusOK, http.StatusText(http.StatusOK))

type errorBase struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Status  string        `json:"status"`
	Detail  []interface{} `json:"detail"`
}
type ResponseBase struct {
	Error errorBase `json:"error"`
}

func NewResponseBase(code int, message string, detail ...interface{}) *ResponseBase {
	return &ResponseBase{
		Error: errorBase{
			Code:    code,
			Message: message,
			Status:  http.StatusText(code),
			Detail:  detail,
		},
	}
}
