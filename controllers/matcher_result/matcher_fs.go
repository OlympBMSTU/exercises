package matcher_result

import (
	"net/http"

	fs "github.com/OlympBMSTU/exercises/fstorage/result"
	"github.com/OlympBMSTU/exercises/result"
)

var mapHttpFsStatuses = map[int]ResultInfo{
	fs.NO_ERROR:          NewResultInfo("Ok", http.StatusOK, statusOK), //, "
	fs.ERROR_CREATE_DIRS: NewResultInfo("Internal server error", http.StatusInternalServerError, statusError),
	fs.ERROR_CREATE_FILE: NewResultInfo("Internal server error", http.StatusInternalServerError, statusError),
	fs.ERROR_WRITE_FILE:  NewResultInfo("Internal server error", http.StatusInternalServerError, statusError),
}

func getAssociatedFsInfo(res result.Result) ResultInfo {
	return mapHttpFsStatuses[res.GetStatus().GetCode()]
}
