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
	UNIQUE_CONSTRAINT = "2305"
	DEFAULT_LIMIT     = 20
	DEFAULT_OFFSET    = 0
)

// todo

// finally works
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

	status := parseError(err)
	if status.IsError() {
		return DbResult{
			DbData{nil},
			status,
		}
	}

	if returnCode == -1 {
		return DbResult{
			DbData{nil},
			DbStatus{
				NO_SUBJECT_ERROR,
				"There is no subject in db",
			},
		}
	}

	status.code = CREATED
	return DbResult{
		DbData{nil},
		status,
	}
}

// var excerciese entities.ExcercieseEntity

// excerciese, err := scanExcerciese(row)
// err := row.Scan(
// 	&excerciese.Id,
// 	&excerciese.AuthorId,
// 	&excerciese.RightAnswer,
// 	&excerciese.Level,
// 	&excerciese.FileName,
// 	&excerciese.Subject,
// )

func GetExcerciese(id uint, pool *pgx.ConnPool) DbResult { //(*views.ExcercieseView, error) {
	rows, err := pool.Query(GET_EXCERCIESE_BY_ID, id)
	defer rows.Close()
	status := parseError(err)
	if status.IsError() {
		return DbResult{
			DbData{nil},
			status,
		}
	}

	excerciese, err := scanExcerciese(rows)

	status = parseError(err)
	if status.IsError() {
		return DbResult{
			DbData{nil},
			status,
		}
	}

	tags, err := getTags(GET_TAGS_FOR_EXCERCIESE, pool, id)
	status = parseError(err)
	if status.IsError() {
		return DbResult{
			DbData{nil},
			status,
		}
	}

	data := CreateDbData(views.ExcercieseViewFrom(*excerciese, *tags))
	return DbResult{
		data,
		status,
	}
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
	status := parseError(err)

	if status.IsError() {
		return DbResult{
			DbData{nil},
			status,
		}
	}

	var entities []entities.ExcercieseEntity
	for rows.Next() {
		// var excerciese entities.ExcercieseEntity
		excerciese, err := scanExcerciese(rows)

		// wtf
		if err != nil {
			break
			// continue
		}
		entities = append(entities, *excerciese)
	}
	// if err != nil {
	status = parseError(err)
	if status.IsError() {
		return DbResult{
			DbData{nil},
			status,
		}
	}
	// }
	return DbResult{
		DbData{entities},
		status,
	}
}

// func GetExcercieseList(tag string, subject string, level int,
// 	limit int, offset int, order_level bool, poll *pgx.ConnPool) (*[]entities.ExcercieseEntity, error) {

// 	var query bytes.Buffer

// 	var args []interface{}
// 	if tag != "" {
// 		query.WriteString(GET_EXCERCIESE_BY_SUBJECT_AND_TAG)
// 		args = append(args, subject, tag, subject)
// 	} else {
// 		args = append(args, subject)
// 		query.WriteString(GET_EXCERCIESES_BY_SUBJECT)
// 	}

// 	if level != -1 {
// 		args = append(args, level)
// 		query.WriteString(fmt.Sprintf("AND ex.level = $%d ", len(args)))
// 	} else {
// 		query.WriteString("ORDER BY ex.level ")
// 		if order_level {
// 			query.WriteString("DESC ")
// 		}
// 	}

// 	if limit == -1 {
// 		limit = DEFAULT_LIMIT
// 	}

// 	args = append(args, limit)
// 	query.WriteString(fmt.Sprintf("LIMIT $%d ", len(args)))

// 	if offset == -1 {
// 		offset = DEFAULT_OFFSET
// 	}

// 	args = append(args, offset)
// 	query.WriteString(fmt.Sprintf("OFFSET $%d ", len(args)))
// 	fmt.Println(query.String())

// 	rows, err := poll.Query(query.String(), args...)

// 	defer rows.Close()
// 	if err != nil {
// 		return nil, err
// 	}

// 	var entities []entities.ExcercieseEntity
// 	for rows.Next() {
// 		// row := (*pgx.Row)(rows)
// 		excerciese, err := scanExcerciese(rows)
// 		// wtf
// 		if err != nil {
// 			return nil, err
// 			// continue
// 		}
// 		entities = append(entities, *excerciese)
// 	}
// 	return &entities, nil
// }
