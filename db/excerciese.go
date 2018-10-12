package db

import (
	"bytes"
	"fmt"

	"github.com/OlympBMSTU/excericieses/entities"
	"github.com/OlympBMSTU/excericieses/views"
	"github.com/jackc/pgx"
	"github.com/lib/pq"
)

// type Error interface {
// 	Value
// }

// type Data interface {
// 	GetData()
// }

// type ScalarData struct {
// 	data interface{}
// 	err Error
// }

// func (scData *ScalarData) GetData() interface{} {
// 	return scData.data
// }

// type Result interface {
// 	Unwrap()
// }

// type RowsResult struct {
// 	data []interface{}
// 	err
// }

const (
	UNIQUE_CONSTRAINT = "2305"
	DEFAULT_LIMIT     = 20
	DEFAULT_OFFSET    = 0
)

// finally works
func SaveExcerciese(excerciese entities.ExcercieseEntity, pool *pgx.ConnPool) error {
	_, err := pool.Exec(insert_excerciese,
		excerciese.GetAuthorId(),
		excerciese.GetRightAnswer(),
		excerciese.GetLevel(),
		excerciese.GetFileName(),
		excerciese.GetSubject(),
		pq.Array(excerciese.GetTags()),
	)

	if err != nil {
		return err
	}
	return nil
}

func scanExcerciese(rows *pgx.Rows) (*entities.ExcercieseEntity, error) {
	var excerciese entities.ExcercieseEntity
	err := rows.Scan(
		&excerciese.Id,
		&excerciese.AuthorId,
		&excerciese.RightAnswer,
		&excerciese.Level,
		&excerciese.FileName,
		&excerciese.Subject,
	)
	if err != nil {
		return nil, err
	}
	return &excerciese, nil
}

func getTags(query string, pool *pgx.ConnPool, args ...interface{}) (*[]string, error) {
	rows, err := pool.Query(query, args[0])
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	var tags []string

	for rows.Next() {
		var tag string
		err := rows.Scan(&tag)
		if err != nil {
			continue
		}
		tags = append(tags, tag)
	}

	return &tags, nil

}

func GetExcerciese(id uint, pool *pgx.ConnPool) (*views.ExcercieseView, error) {
	var excerciese entities.ExcercieseEntity
	row := pool.QueryRow(get_excerciese, id)

	// excerciese, err := scanExcerciese(row)
	err := row.Scan(
		&excerciese.Id,
		&excerciese.AuthorId,
		&excerciese.RightAnswer,
		&excerciese.Level,
		&excerciese.FileName,
		&excerciese.Subject,
	)

	// todo parse error to my errors
	// or map error

	if err != nil {
		return nil, err
	}

	tags, err := getTags(get_tags_for_excerciese, pool, id)
	if err != nil {
		return nil, err
	}

	view := views.ExcercieseViewFrom(excerciese, *tags)
	return &view, nil
}

func GetExcercieseList(tag string, subject string, level int,
	limit int, offset int, order_level bool, poll *pgx.ConnPool) (*[]entities.ExcercieseEntity, error) {
	var query bytes.Buffer

	var args []interface{}
	if tag != "" {
		query.WriteString(GET_EXCERCIESE_BY_SUBJECT_AND_TAG)
		args = append(args, subject, tag, subject)
	} else {
		args = append(args, subject)
		query.WriteString(GET_BY_SUBJECT)
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
	fmt.Println(query.String())

	rows, err := poll.Query(query.String(), args...)

	defer rows.Close()
	if err != nil {
		return nil, err
	}

	var entities []entities.ExcercieseEntity
	for rows.Next() {
		// row := (*pgx.Row)(rows)
		excerciese, err := scanExcerciese(rows)
		// wtf
		if err != nil {
			continue
		}
		entities = append(entities, *excerciese)
	}
	return &entities, nil
}

func GetTgasBySubect(subject string, pool *pgx.ConnPool) (*[]string, error) {
	return getTags(get_tags_by_subject, pool, subject)
}
