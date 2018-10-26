package matcher_result

import (
	"fmt"
	"net/http"

	http_res "github.com/OlympBMSTU/excericieses/controllers/http_result"
	db "github.com/OlympBMSTU/excericieses/db/result"
	"github.com/OlympBMSTU/excericieses/result"
)

// json: {
// "status": Ok,
// "message": error,
// "data": data
// }

var mapHttpDbStatuses = map[int]http_res.HttpResult{
	db.NO_ERROR:         http_res.CreateHttpResult(http.StatusOK, "OK"),
	db.CREATED:          http_res.CreateHttpResult(http.StatusCreated, "Created"),
	db.QUERY_ERROR:      http_res.ResultNotFound(),
	db.EMPTY_RESULT:     http_res.ResultNotFound(),
	db.DB_CONN_ERROR:    http_res.ResultInernalSreverError(),
	db.PARSE_ERROR:      http_res.ResultInernalSreverError(),
	db.CONSTRAINT_ERROR: http_res.CreateHttpResult(http.StatusConflict, "Conflict"),
	db.NO_SUBJECT_ERROR: http_res.CreateHttpResult(http.StatusBadRequest, ""),
}

func MatchDbResult(res result.Result) http_res.HttpResult {
	status := res.GetStatus()
	outHttpRes := mapHttpDbStatuses[status.GetCode()]
	// if outHttpRes == nil {

	// }
	fmt.Print(outHttpRes)

	return http_res.ResultInernalSreverError()
	// if status.IsError() {

	// }

}
