package matcher_result

import (
	"net/http"

	auth_res "github.com/OlympBMSTU/exercises/auth/result"
	"github.com/OlympBMSTU/exercises/result"
)

var mapHttpAuthStatuses = map[int]ResultInfo{
	auth_res.NO_ERROR:        NewResultInfo("Ok", http.StatusOK, statusOK),
	auth_res.NO_AUTHROIZED:   NewResultInfo("Unauthorized", http.StatusUnauthorized, statusError),
	auth_res.ERROR_PARSE_JWT: NewResultInfo("Unauthorized", http.StatusUnauthorized, statusError),
	auth_res.NO_COOKIE:       NewResultInfo("Unauthorized", http.StatusUnauthorized, statusError),
}

func getAssociatedAuthInfo(res result.Result) ResultInfo {
	return mapHttpAuthStatuses[res.GetStatus().GetCode()]
}
