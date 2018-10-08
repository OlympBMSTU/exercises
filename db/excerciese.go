package db

import (
	"fmt"

	"github.com/OlympBMSTU/excericieses/entities"
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

func scanExcerciese(row *pgx.Rows) (*entities.ExcercieseEntity, error) {
	var excerciese entities.ExcercieseEntity
	err := row.Scan(
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

func GetExcerciese(id uint, poll *pgx.ConnPool) (*entities.ExcercieseEntity, error) {
	var excerciese entities.ExcercieseEntity
	row := poll.QueryRow(get_excerciese, id)

	err := row.Scan(
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

func GetExcercieseList(tag string, subject string, level int,
	limit int, offset int, order_level bool, poll *pgx.ConnPool) (*[]entities.ExcercieseEntity, error) {
	query := ""

	if tag != "" {
		query = fmt.Sprintf(get_by_subject_tag, subject, tag, subject)
	} else {
		query = fmt.Sprintf(get_by_subject, subject)
	}

	if level != -1 {
		query += fmt.Sprintf("AND ex.level = %d ", level)
	} else {
		query += "ORDER BY ex.level "
		if order_level {
			query += "DESC "
		}
	}

	if limit == -1 {
		limit = DEFAULT_LIMIT
	}

	query += fmt.Sprintf("LIMIT %d ", limit)

	if offset == -1 {
		offset = DEFAULT_OFFSET
	}

	query += fmt.Sprintf("OFFSET %d ", offset)

	// var rows pgx.Rows
	// var err error
	// if tag == "" {
	// 	rows, err := poll.Query(query, subject)
	// } else {
	// 	rows, err := poll.Query(query, subject, tag, subject)
	// }
	// // fmt.Println(query)
	fmt.Println(query)
	rows, err := poll.Query(query)
	// rows, err := poll.Query(query, subject, tag, subject)

	if err != nil {
		return nil, err
	}

	// todo make array

	entities := make([]entities.ExcercieseEntity, 10)

	idx := 0
	for rows.Next() {
		excerciese, err := scanExcerciese(rows)
		// wtf
		if err != nil {
			continue
		}
		entities[idx] = *excerciese
		idx += 1
	}
	return &entities, nil
}
