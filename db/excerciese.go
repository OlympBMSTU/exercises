package db

import (
	"github.com/OlympBMSTU/excericieses/entities"
	"github.com/jackc/pgx"
	"github.com/lib/pq"
)

const UNIQUE_CONSTRAINT = "2305"

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

func GetExcercieseList(tags string, subject, level uint, limit uint, offset uint, order_level bool) error {
	return nil
}
