package matcher_result

import (
	auth_result "github.com/OlympBMSTU/excericieses/auth/result"
	http "github.com/OlympBMSTU/excericieses/controllers/http_result"
	db_result "github.com/OlympBMSTU/excericieses/db/result"
	fs_result "github.com/OlympBMSTU/excericieses/fstorage/result"
	root_result "github.com/OlympBMSTU/excericieses/result"
)

// func MatchOkResult() {

// }

// func MatchErrorResult(res root_result.Result) http.HttpResult {
// 	switch res.(type) {
// 	case db_result.DbResult:
// 		return MatchDbResult(res)
// 	case fs_result.FSResult:
// 		return MatchFSResult(res)
// 	case auth_result.AuthResult:
// 		return MatchAuthResult(res)
// 	default:
// 		return http.ResultInernalSreverError()
// 	}
// }

func MatchResult(res root_result.Result) http.HttpResult {
	// if !res.IsError() {
	// 	code := 200
	// 	status := "Ok"
	// 	message := "Ok"

	// 	jsonRes := output.ResultView{}
	// 	jsonRes.SetData(res.GetData())
	// 	jsonRes.SetMessage(message)
	// 	jsonRes.SetStatus(status)
	// 	body, err := json.Marshal(jsonRes)
	// 	//
	// 	if err != nil {
	// 		return http.ResultInernalSreverError()
	// 	}
	// 	fmt.Print(err)

	// 	return http.NewHttpResult(code, body)
	// }

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
