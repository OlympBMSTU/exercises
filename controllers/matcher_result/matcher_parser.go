package matcher_result

import (
	"net/http"

	parser "github.com/OlympBMSTU/exercises/parser/result"
	"github.com/OlympBMSTU/exercises/result"
)

var mapHttpParserStatuses = map[int]ResultInfo{
	parser.NO_ERROR:        NewResultInfo("Ok", http.StatusOK, statusOK),
	parser.INCORRECT_BODY:  NewResultInfo("Incorrect body", http.StatusBadRequest, statusError),
	parser.INCORRECT_LEVEL: NewResultInfo("Incorrect level", http.StatusBadRequest, statusError),
	parser.INCORRECT_TAGS:  NewResultInfo("Incorrect tags", http.StatusBadRequest, statusError),
}

func getAssociatedParserInfo(res result.Result) ResultInfo {
	return mapHttpParserStatuses[res.GetStatus().GetCode()]
}
