package db

import (
	"github.com/OlympBMSTU/exercises/db/result"
	"github.com/jackc/pgx"
)

func AddSubject(subject string, pool *pgx.ConnPool) result.DbResult {
	_, err := pool.Exec(ADD_SUBJECT, subject)
	return result.CreateResult(nil, err, result.CREATED)
}

func GetSubjects(pool *pgx.ConnPool) result.DbResult {
	rows, err := pool.Query(GET_SUBJECTS)
	if err != nil {
		return result.ErrorResult(err)
	}

	var subjects []string
	for rows.Next() {
		var subject string

		err = rows.Scan(&subject)
		if err != nil {
			break
		}

		subjects = append(subjects, subject)
	}

	if err != nil {
		return result.ErrorResult(err)
	}

	if len(subjects) == 0 {
		return result.ErrorResult(result.EMPTY_RESULT, "")
	}

	return result.OkResult(subjects)
}

func GetTgasBySubect(subject string, pool *pgx.ConnPool) result.DbResult {
	data, err := getTags(GET_TAGS_BY_SUBJECT, pool, subject)
	if err != nil {
		return result.ErrorResult(err)
	}

	if len(*data) == 0 {
		return result.ErrorResult(result.EMPTY_RESULT, "")
	}

	return result.OkResult(data)
}
