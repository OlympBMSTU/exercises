package db

import (
	"github.com/HustonMmmavr/excercieses/entities"
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
	rows, err := poll.Query(get_excerciese, id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// its fucking crutch
	idx := 0
	for rows.Next() {
		idx += 1
		rows.Scan(
			&excerciese.Id,
			&excerciese.AuthorId,
			&excerciese.RightAnswer,
			&excerciese.Level,
			&excerciese.FileName,
			&excerciese.Subject,
		)
	}

	if idx == 0 {
		return nil, nil
	}
	return &excerciese, nil
}

func GetExcercieseList() {

}
