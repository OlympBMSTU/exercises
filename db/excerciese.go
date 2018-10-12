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

// finally works
func SaveExcerciese(excerciese entities.ExcercieseEntity, pool *pgx.ConnPool) error {
	_, err := pool.Exec(INSERT_EXCERCIESE,
		excerciese.GetAuthorId(),
		excerciese.GetRightAnswer(),
		excerciese.GetLevel(),
		excerciese.GetFileName(),
		excerciese.GetSubject(),
		pq.Array(excerciese.GetTags()),
	)

	// parseError(err)

	if err != nil {
		return err
	}

	// if res == -1 {
	// 	return nil, NO_SUBJECT_ERROR
	// }

	return nil
}

func GetExcerciese(id uint, pool *pgx.ConnPool) (*views.ExcercieseView, error) {
	var excerciese entities.ExcercieseEntity
	row := pool.QueryRow(GET_EXCERCIESE_BY_ID, id)

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

	tags, err := getTags(GET_TAGS_FOR_EXCERCIESE, pool, id)
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
