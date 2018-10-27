package matcher_result

import (
	"encoding/json"
	"net/http"

	http_res "github.com/OlympBMSTU/exercises/controllers/http_result"
	fs "github.com/OlympBMSTU/exercises/fstorage/result"
	"github.com/OlympBMSTU/exercises/result"
	"github.com/OlympBMSTU/exercises/views/output"
)

var mapHttpFsStatuses = map[int]ResultInfo{
	fs.NO_ERROR:          NewResultInfo("Ok", http.StatusOK, statusOK), //, "
	fs.ERROR_CREATE_DIRS: NewResultInfo("Internal server error", http.StatusInternalServerError, statusError),
	fs.ERROR_CREATE_FILE: NewResultInfo("Internal server error", http.StatusInternalServerError, statusError),
	fs.ERROR_WRITE_FILE:  NewResultInfo("Internal server error", http.StatusInternalServerError, statusError),
}

func MatchFSResult(res result.Result) http_res.HttpResult {
	var jsonRes output.ResultView
	info := mapHttpDbStatuses[res.GetStatus().GetCode()]

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
