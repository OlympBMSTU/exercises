package matcher_result

import (
	"net/http"

	"github.com/OlympBMSTU/exercises/result"
	sender "github.com/OlympBMSTU/exercises/sender/result"
)

const (
	NO_ERROR          = 0
	ERROR_SEND        = 1
	ERROR_CREATE_JSON = 2
)

var mapHttpSenderStatuses = map[int]ResultInfo{
	sender.NO_ERROR:          NewResultInfo("Ok", http.StatusOK, statusOK),
	sender.ERROR_SEND:        NewResultInfo("Internal server error send", http.StatusInternalServerError, statusError),
	sender.ERROR_CREATE_JSON: NewResultInfo("Internal server error send", http.StatusInternalServerError, statusError),
}

func getAssociatedSenderInfo(res result.Result) ResultInfo {
	return mapHttpSenderStatuses[res.GetStatus().GetCode()]
}
