package db

import (
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

func getTags(query string, pool *pgx.ConnPool, args ...interface{}) (*[]string, error) {
	rows, err := pool.Query(query, args[0])
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

	tags, err := getTags(get_tags_for_excerciese, pool, id)
	if err != nil {
		return nil, err
	}

	view := views.ExcercieseViewFrom(excerciese, *tags)
	return &view, nil
}

func GetExcercieseList(tag string, subject string, level int,
	limit int, offset int, order_level bool, poll *pgx.ConnPool) (*[]entities.ExcercieseEntity, error) {
	query := ""
	var args []interface{}
	idx_args := 4
	if tag != "" {
		query = GET_EXCERCIESE_BY_SUBJECT_AND_TAG //fmt.Sprintf(//get_by_subject_tag, subject, tag, subject)
		args = append(args, subject, tag, subject)
	} else {
		query = fmt.Sprintf(get_by_subject, subject)
	}

	if level != -1 {
		query += fmt.Sprintf("AND ex.level = $%d ", idx_args)
		idx_args += 1
	} else {
		query += "ORDER BY ex.level "
		if order_level {
			query += "DESC "
		}
	}

	if limit == -1 {
		limit = DEFAULT_LIMIT
	}

	query += fmt.Sprintf("LIMIT $%d ", idx_args) //limit)
	idx_args += 1

	if offset == -1 {
		offset = DEFAULT_OFFSET
	}

	query += fmt.Sprintf("OFFSET $%d ", idx_args) //offset)
	fmt.Println(query)

	rows, err := poll.Query(query, args)

	if err != nil {
		return nil, err
	}

	var entities []entities.ExcercieseEntity
	for rows.Next() {
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
	// rows, err := pool.Query(get_tags_by_subject, subject)

	// if err != nil {
	// 	return nil, err
	// }

	tags, err := getTags(get_tags_by_subject, pool, subject)

	return tags, err

}

// var rows pgx.Rows
// var err error
// if tag == "" {
// 	rows, err := poll.Query(query, subject)
// } else {
// 	rows, err := poll.Query(query, subject, tag, subject)
// }
// // fmt.Println(query)

// rows, err := poll.Query(query, subject, tag, subject)
