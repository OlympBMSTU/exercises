package matcher_result

import (
	"encoding/json"
	"net/http"

	http_res "github.com/OlympBMSTU/excericieses/controllers/http_result"
	db "github.com/OlympBMSTU/excericieses/db/result"
	"github.com/OlympBMSTU/excericieses/result"
	"github.com/OlympBMSTU/excericieses/views/output"
)

var mapHttpDbStatuses = map[int]ResultInfo{
	db.NO_ERROR:         NewResultInfo("Ok", http.StatusOK, statusOK), //, "OK"),
	db.CREATED:          NewResultInfo("Created", http.StatusCreated, statusOK),
	db.QUERY_ERROR:      NewResultInfo("Not found", http.StatusNotFound, statusError),
	db.EMPTY_RESULT:     NewResultInfo("Not found", http.StatusNotFound, statusError),
	db.DB_CONN_ERROR:    NewResultInfo("Internal server error", http.StatusInternalServerError, statusError),
	db.PARSE_ERROR:      NewResultInfo("Internal server error", http.StatusInternalServerError, statusError),
	db.CONSTRAINT_ERROR: NewResultInfo("Connflict", http.StatusConflict, statusError),
	db.NO_SUBJECT_ERROR: NewResultInfo("Please use correct subject", http.StatusBadRequest, statusSubjectError),
}

func MatchDbResult(res result.Result) http_res.HttpResult {
	var jsonRes output.ResultView
	info := mapHttpDbStatuses[res.GetStatus().GetCode()]
	if res.IsError() {
		jsonRes.SetData(nil)
	} else {
		jsonRes.SetData(res.GetData())
	}
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
