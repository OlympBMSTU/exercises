package matcher_result

import (
	"encoding/json"
	"net/http"

	auth_res "github.com/OlympBMSTU/exercises/auth/result"
	http_res "github.com/OlympBMSTU/exercises/controllers/http_result"
	"github.com/OlympBMSTU/exercises/result"
	"github.com/OlympBMSTU/exercises/views/output"
)

var mapHttpAuthStatuses = map[int]ResultInfo{
	auth_res.NO_ERROR:        NewResultInfo("Ok", http.StatusOK, statusOK),
	auth_res.NO_AUTHROIZED:   NewResultInfo("Unauthorized", http.StatusUnauthorized, statusError),
	auth_res.ERROR_PARSE_JWT: NewResultInfo("Unauthorized", http.StatusUnauthorized, statusError),
	auth_res.NO_COOKIE:       NewResultInfo("Unauthorized", http.StatusUnauthorized, statusError),
}

func MatchAuthResult(res result.Result) http_res.HttpResult {
	var jsonRes output.ResultView
	info := mapHttpAuthStatuses[res.GetStatus().GetCode()]
	jsonRes.SetStatus(info.Status)
	jsonRes.SetMessage(info.Message)
	code := info.HttpCode

	var outHttpRes http_res.HttpResult
	val, err := json.Marshal(jsonRes)
	if err != nil {
		code = http.StatusInternalServerError
	} else {
		outHttpRes.SetBody(val)
	}
	outHttpRes.SetStatus(code)

	return outHttpRes
}
