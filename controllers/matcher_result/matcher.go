package matcher_result

import (
	auth_result "github.com/OlympBMSTU/excericieses/auth/result"
	db_result "github.com/OlympBMSTU/excericieses/db/result"
	fs_result "github.com/OlympBMSTU/excericieses/fstorage/result"
	root_result "github.com/OlympBMSTU/excericieses/result"
)

func MatchResult(res root_result.Result) {
	switch res.(type) {
	case db_result.DbResult:
		return MatchDbResult(res)
	case fs_result.FSResult:
		return MatchFSResult(res)
	case auth_result.AuthResult:
		return MatchAuthResult(res)
	default:
		return
	}
}
