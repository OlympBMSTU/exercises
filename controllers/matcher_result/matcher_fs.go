package matcher_result

import (
	"fmt"
	"net/http"

	http_res "github.com/OlympBMSTU/excericieses/controllers/http_result"
	fs "github.com/OlympBMSTU/excericieses/fstorage/result"
	"github.com/OlympBMSTU/excericieses/result"
)

var mapHttpFsStatuses = map[int]http_res.HttpResult{
	fs.NO_ERROR:          http_res.CreateHttpResult(http.StatusOK, "OK"),
	fs.ERROR_CREATE_DIRS: http_res.ResultInernalSreverError(),
	fs.ERROR_CREATE_FILE: http_res.ResultInernalSreverError(),
	fs.ERROR_WRITE_FILE:  http_res.ResultInernalSreverError(),
}

func MatchFSResult(res result.Result) http_res.HttpResult {
	var code int
	status := res.GetStatus()
	switch status.GetCode() {
	case fs.NO_ERROR:
		code = http.StatusOK
	default:
		code = http.StatusInternalServerError
	}
	fmt.Print(code)
	return http_res.ResultInernalSreverError()
}
