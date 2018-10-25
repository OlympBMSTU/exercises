package db

import (
	"github.com/OlympBMSTU/excericieses/entities"
	"github.com/jackc/pgx"
)

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
