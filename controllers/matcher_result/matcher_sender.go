package matcher_result

import (
	http_res "github.com/OlympBMSTU/exercises/controllers/http_result"
	"github.com/OlympBMSTU/exercises/result"
)

func MatchSenderResult(res result.Result) http_res.HttpResult {
	return http_res.ResultInernalSreverError()
}
