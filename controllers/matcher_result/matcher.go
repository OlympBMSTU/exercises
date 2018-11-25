package matcher_result

import (
	"encoding/json"
	"net/http"

	auth_result "github.com/OlympBMSTU/exercises/auth/result"
	http_res "github.com/OlympBMSTU/exercises/controllers/http_result"
	db_result "github.com/OlympBMSTU/exercises/db/result"
	fs_result "github.com/OlympBMSTU/exercises/fstorage/result"
	parser_result "github.com/OlympBMSTU/exercises/parser/result"
	root_result "github.com/OlympBMSTU/exercises/result"
	sender_result "github.com/OlympBMSTU/exercises/sender/result"
	"github.com/OlympBMSTU/exercises/views/output"
)

func fillResult(info ResultInfo, body interface{}) http_res.HttpResult {
	var jsonRes output.ResultView
	jsonRes.SetData(body)
	jsonRes.SetStatus(info.Status)
	jsonRes.SetMessage(info.Message)

	val, err := json.Marshal(jsonRes)
	code := info.HttpCode
	var outHttpRes http_res.HttpResult

	if err != nil {
		code = http.StatusInternalServerError
	} else {
		outHttpRes.SetBody(val)
	}

	outHttpRes.SetStatus(code)
	return outHttpRes
}

// func MatchDbResult(res result.Result) http_res.HttpResult {
// 	var jsonRes output.ResultView
// 	info := mapHttpDbStatuses[res.GetStatus().GetCode()]
// 	if res.IsError() {
// 		jsonRes.SetData(nil)
// 	} else {
// 		jsonRes.SetData(res.GetData())
// 	}
// 	jsonRes.SetStatus(info.Status)
// 	jsonRes.SetMessage(info.Message)

// 	val, err := json.Marshal(jsonRes)
// 	code := info.HttpCode
// 	var outHttpRes http_res.HttpResult

// 	if err != nil {
// 		code = http.StatusInternalServerError
// 	} else {
// 		outHttpRes.SetBody(val)
// 	}

// 	outHttpRes.SetStatus(code)
// 	return outHttpRes
// }

func MatchResult(res root_result.Result) http_res.HttpResult {
	var infoRes ResultInfo
	var bodyData interface{}
	bodyData = nil

	switch res.(type) {
	case db_result.DbResult:
		if !res.IsError() {
			bodyData = res.GetData()
		}
		infoRes = getAssociatedDbInfo(res)
		// return MatchDbResult(res)
	case fs_result.FSResult:
		infoRes = getAssociatedFsInfo(res)
		// return MatchFSResult(res)
	case auth_result.AuthResult:
		infoRes = getAssociatedAuthInfo(res)
		// return MatchAuthResult(res)
	case sender_result.SenderResult:
		infoRes = getAssociatedSenderInfo(res)
		// return MatchSenderResult(res)
	case parser_result.ParserResult:
		infoRes = getAssociatedParserInfo(res)
		infoRes.Message += res.GetStatus().GetDescription()
		// return MatchParserResult(res)
	default:
		// coorect this
		return http_res.ResultInernalSreverError()
	}

	return fillResult(infoRes, bodyData)
}
