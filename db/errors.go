package db

import (
	"github.com/jackc/pgx"
)

// value, string for http
const (
	NO_ERROR         = 0 // Ok
	CREATED          = 1 // Ok
	PARSE_ERROR      = 2 // Same that empty result, maybe need to delete, also or db error,
	QUERY_ERROR      = 3 // Its unused also dberror
	DB_CONN_ERROR    = 4 // No db, table, connection dropped
	EMPTY_RESULT     = 5 // No data for query
	CONSTRAINT_ERROR = 6 // Duplicates unique keys --- not used
	NO_SUBJECT_ERROR = 7 // While inserting new excerciese we have to choose exisiting subject
)

const PG_UNIQUE_CONSTRAINT_CODE = "23505"
const FUNDAMENTAL_QUERY_EMPTY_ERROR = "no rows in result set"

// 1-2 ok
// 3,5 is one error
// 6 ?
// 7 pg error UNIQUE_CONSTRAINT
// 8 custom my error
type DbData struct {
	data interface{}
}

func CreateDbData(data interface{}) DbData {
	return DbData{
		data,
	}
}

func (dat *DbData) GetData() interface{} {
	return dat.data
}

type DbResult struct {
	data DbData
	err  DbStatus
}

func (res *DbResult) SetResult(data DbData) {
	res.data = data
}

func (res *DbResult) SetError(err DbStatus) {
	res.err = err
}

// func (res *DbR)
//
type DbStatus struct {
	code  int
	descr string
}

func DefaultStatus() DbStatus {
	return DbStatus{
		NO_ERROR,
		"",
	}
}

func (status *DbStatus) SetCode(code int) {
	status.code = code
}

func (status *DbStatus) IsError() bool {
	return status.code == NO_ERROR || status.code == CREATED
}

func parseError(err error) DbStatus {
	var code int
	var descr string
	if err == nil {
		code = NO_ERROR
		descr = ""
	} else {
		switch typedError := err.(type) {
		case pgx.PgError:
			if typedError.Code == PG_UNIQUE_CONSTRAINT_CODE {
				code = CONSTRAINT_ERROR
			} else {
				// uncategorized error
				code = DB_CONN_ERROR
			}
			descr = typedError.Message

		// other errors, also fundamental from pkg/errors
		default:
			// its fucking crutch, but pgx returns fundamental error
			// which is private when row.Scan fails
			// fundamental - is a private package type of package pkg/error
			// also fundamental returns, when there is not enough sended
			// parameters to query
			if typedError.Error() == FUNDAMENTAL_QUERY_EMPTY_ERROR {
				code = QUERY_ERROR
			} else {
				code = DB_CONN_ERROR
			}
			descr = typedError.Error()

		}
	}

	return DbStatus{
		code:  code,
		descr: descr,
	}
}
