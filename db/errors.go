package db

import (
	"fmt"

	"github.com/jackc/pgx"
)

// value, string for http
const (
	NO_ERROR         = 0 // Ok
	CREATED          = 1 // Ok
	PARSE_ERROR      = 2 // Same that empty result, maybe need to delete, also
	QUERY_ERROR      = 3 // Its unused
	DB_CONN_ERROR    = 4 // No db, table, connection dropped
	EMPTY_RESULT     = 5 // No data for query
	CONSTRAINT_ERROR = 6 // Duplicates unique keys
	NO_SUBJECT_ERROR = 7 // While inserting new excerciese we have to choose exisiting subject
)

type DbData struct {
	data interface{}
}

// func (dat)

func (dat *DbData) GetData() interface{} {
	return dat.data
}

type DbResult struct {
	data DbData
	err  DbError
}

func (res *DbResult) SetResult(data DbData) {
	res.data = data
}

func (res *DbResult) SetError(err DbError) {
	res.err = err
}

// func (res *DbR)
//
type DbError struct {
	code  int
	descr string
}

func parseError(err error) DbError {
	var code int
	var descr string
	if err == nil {
		code = NO_ERROR
		descr = ""
	} else {
		switch t := err.(type) {
		case *pgx.PgError:
			fmt.Println("DB_CONN_ERR", err)
		default:
			fmt.Println("def", err, t)
			// case: *pkg.fundamential: fmt.Println("funcdamential ", err)
		}
	}

	return DbError{
		code:  code,
		descr: descr,
	}
}
