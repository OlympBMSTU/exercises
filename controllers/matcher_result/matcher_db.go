package matcher_result

import (
	"net/http"

	db "github.com/OlympBMSTU/exercises/db/result"
	"github.com/OlympBMSTU/exercises/result"
)

var mapHttpDbStatuses = map[int]ResultInfo{
	db.NO_ERROR:     NewResultInfo("Ok", http.StatusOK, statusOK), //, "OK"),
	db.CREATED:      NewResultInfo("Created", http.StatusCreated, statusOK),
	db.QUERY_ERROR:  NewResultInfo("Ok", http.StatusOK, statusOK),
	db.EMPTY_RESULT: NewResultInfo("Ok", http.StatusOK, statusOK),
	// db.QUERY_ERROR:      NewResultInfo("Not found", http.StatusNotFound, statusError),
	// db.EMPTY_RESULT:     NewResultInfo("Not found", http.StatusNotFound, statusError),
	db.DB_CONN_ERROR:    NewResultInfo("Internal server error", http.StatusInternalServerError, statusError),
	db.PARSE_ERROR:      NewResultInfo("Internal server error", http.StatusInternalServerError, statusError),
	db.CONSTRAINT_ERROR: NewResultInfo("Connflict", http.StatusConflict, statusError),
	db.NO_SUBJECT_ERROR: NewResultInfo("Please use correct subject", http.StatusBadRequest, statusSubjectError),
}

func getAssociatedDbInfo(res result.Result) ResultInfo {
	return mapHttpDbStatuses[res.GetStatus().GetCode()]
}
