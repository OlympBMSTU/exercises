package db

import (
	"bytes"
	"fmt"

	"github.com/OlympBMSTU/excericieses/entities"
	"github.com/OlympBMSTU/excericieses/views"
	"github.com/jackc/pgx"
	"github.com/lib/pq"
)

const (
	DEFAULT_LIMIT  = 20
	DEFAULT_OFFSET = 0
)

func SaveExcerciese(excerciese entities.ExcercieseEntity, pool *pgx.ConnPool) DbResult {
	row := pool.QueryRow(INSERT_EXCERCIESE,
		excerciese.GetAuthorId(),
		excerciese.GetRightAnswer(),
		excerciese.GetLevel(),
		excerciese.GetFileName(),
		excerciese.GetSubject(),
		pq.Array(excerciese.GetTags()),
	)

	var returnCode int
	err := row.Scan(&returnCode)

	if err != nil {
		return errorResult(err)
	}

	if returnCode == -1 {
		return errorResult(NO_SUBJECT_ERROR, "There is no subject in db")
	}

	return okResult(nil, CREATED)
}

func GetExcerciese(id uint, pool *pgx.ConnPool) DbResult {
	rows, err := pool.Query(GET_EXCERCIESE_BY_ID, id)
	defer rows.Close()

	if err != nil {
		return errorResult(err)
	}

	// todo scan ex returns also dbres
	var excerciese *entities.ExcercieseEntity
	for rows.Next() {
		excerciese, err = scanExcerciese(rows)
	}

	if err != nil {
		return errorResult(err)
	}

	if excerciese == nil {
		return errorResult(EMPTY_RESULT, "")
	}

	tags, err := getTags(GET_TAGS_FOR_EXCERCIESE, pool, id)
	if err != nil {
		return errorResult(err)
	}

	return okResult(views.ExcercieseViewFrom(*excerciese, *tags))
}

func GetExcercieseList(tag string, subject string, level int,
	limit int, offset int, order_level bool, poll *pgx.ConnPool) DbResult {

	var query bytes.Buffer

	var args []interface{}
	if tag != "" {
		query.WriteString(GET_EXCERCIESE_BY_SUBJECT_AND_TAG)
		args = append(args, subject, tag, subject)
	} else {
		args = append(args, subject)
		query.WriteString(GET_EXCERCIESES_BY_SUBJECT)
	}

	if level != -1 {
		args = append(args, level)
		query.WriteString(fmt.Sprintf("AND ex.level = $%d ", len(args)))
	} else {
		query.WriteString("ORDER BY ex.level ")
		if order_level {
			query.WriteString("DESC ")
		}
	}

	if limit == -1 {
		limit = DEFAULT_LIMIT
	}

	args = append(args, limit)
	query.WriteString(fmt.Sprintf("LIMIT $%d ", len(args)))

	if offset == -1 {
		offset = DEFAULT_OFFSET
	}

	args = append(args, offset)
	query.WriteString(fmt.Sprintf("OFFSET $%d ", len(args)))

	rows, err := poll.Query(query.String(), args...)
	defer rows.Close()

	if err != nil {
		return errorResult(err)
	}

	var entities []entities.ExcercieseEntity
	for rows.Next() {
		excerciese, err := scanExcerciese(rows)

		if err != nil {
			return errorResult(err)
		}
		entities = append(entities, *excerciese)
	}

	return okResult(entities)
}

// status := parseError(err)
// if status.IsError() {
// 	return DbResult{
// 		DbData{nil},
// 		status,
// 	}
// }

// status = parseError(err)
// if status.IsError() {
// 	return DbResult{
// 		DbData{nil},
// 		status,
// 	}
// }

// return DbResult{
// 	DbData{entities},
// 	status,
// }
// data := CreateDbData(views.ExcercieseViewFrom(*excerciese, *tags))

// return DbResult{
// 	DbData{nil},
// 	DbStatus{EMPTY_RESULT, ""},
// }
