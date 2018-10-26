package http_result

import (
	"net/http"
)

type HttpResult struct {
	body   []byte
	status int
}

func NewHttpResult(code int, body []byte) HttpResult {
	return HttpResult{
		body:   body,
		status: code,
	}
}

func (http_res *HttpResult) SetStatus(status int) {
	http_res.status = status
}

func (http_res HttpResult) GetStatus() int {
	return http_res.status
}

func (http_res *HttpResult) SetBody(body []byte) {
	http_res.body = body
}

func (http_res HttpResult) GetData() []byte {
	return http_res.body
}

func CreateHttpResult(code int, message string) HttpResult {
	return HttpResult{
		// data:    nil,
		// message: message,
		status: code,
	}
}

func ResultInernalSreverError() HttpResult {
	return CreateHttpResult(http.StatusInternalServerError, "Internal srever error")
}

func ResultNotFound() HttpResult {
	return CreateHttpResult(http.StatusNotFound, "Not found")
}
