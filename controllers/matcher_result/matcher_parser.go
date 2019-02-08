package matcher_result

import (
	"net/http"

	parser "github.com/OlympBMSTU/exercises/parser/result"
	"github.com/OlympBMSTU/exercises/result"
)

var mapHttpParserStatuses = map[int]ResultInfo{
	parser.NO_ERROR:                  NewResultInfo("Ok", http.StatusOK, statusOK),
	parser.INCORRECT_BODY:            NewResultInfo("Incorrect body", http.StatusBadRequest, statusError),
	parser.INCORRECT_LEVEL:           NewResultInfo("Incorrect level", http.StatusBadRequest, statusError),
	parser.INCORRECT_TAGS:            NewResultInfo("Incorrect tags: ", http.StatusBadRequest, statusError),
	parser.INCORRECT_TYPE_OLYMP:      NewResultInfo("Incorrect type olymp", http.StatusBadRequest, statusError),
	parser.INCORRECT_CLASS:           NewResultInfo("Incorrect class", http.StatusBadRequest, statusError),
	parser.INCORRECT_MARK:            NewResultInfo("Incorrect mark", http.StatusBadRequest, statusError),
	parser.INCORRECT_POSITION:        NewResultInfo("Incorrect position", http.StatusBadRequest, statusError),
	parser.INCORRECT_ANSWER_ARR:      NewResultInfo("Incorrect answers", http.StatusBadRequest, statusError),
	parser.INCORRECT_MARK_IN_ANSWERS: NewResultInfo("Mark in exersize doesn't equals sum marks in answers", http.StatusBadRequest, statusError),
}

func getAssociatedParserInfo(res result.Result) ResultInfo {
	return mapHttpParserStatuses[res.GetStatus().GetCode()]
}
