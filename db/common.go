package db

import (
	"github.com/OlympBMSTU/excericieses/entities"
	"github.com/jackc/pgx"
)

func scanExercise(rows *pgx.Rows) (*entities.ExerciseEntity, error) {
	var exercise entities.ExerciseEntity

	err := rows.Scan(
		&exercise.Id,
		&exercise.AuthorId,
		&exercise.RightAnswer,
		&exercise.Level,
		&exercise.FileName,
		&exercise.Subject,
	)

	if err != nil {
		return nil, err
	}

	return &exercise, nil
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
