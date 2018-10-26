package matcher_result

import (
	"fmt"
	"net/http"

	auth_res "github.com/OlympBMSTU/excericieses/auth/result"
	http_res "github.com/OlympBMSTU/excericieses/controllers/http_result"
	"github.com/OlympBMSTU/excericieses/result"
)

// const fs_http_codes = map[int]

func MatchAuthResult(res result.Result) http_res.HttpResult {
	status := res.GetStatus()
	var code int
	switch status.GetCode() {
	case auth_res.NO_ERROR:
		code = http.StatusOK
	default:
		code = http.StatusUnauthorized
	}
	fmt.Print(code)

	return http_res.ResultInernalSreverError()
}
