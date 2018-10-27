package matcher_result

import (
	auth_result "github.com/OlympBMSTU/exercises/auth/result"
	http "github.com/OlympBMSTU/exercises/controllers/http_result"
	db_result "github.com/OlympBMSTU/exercises/db/result"
	fs_result "github.com/OlympBMSTU/exercises/fstorage/result"
	root_result "github.com/OlympBMSTU/exercises/result"
)

func MatchResult(res root_result.Result) http.HttpResult {
	switch res.(type) {
	case db_result.DbResult:
		return MatchDbResult(res)
	case fs_result.FSResult:
		return MatchFSResult(res)
	case auth_result.AuthResult:
		return MatchAuthResult(res)
	default:
		// coorect this
		return http.ResultInernalSreverError()
	}
}
