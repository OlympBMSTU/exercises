package db

import "github.com/jackc/pgx"

func AddSubject(subject string, pool *pgx.ConnPool) DbResult {
	_, err := pool.Exec(ADD_SUBJECT, subject)
	return DbResult{
		DbData{nil},
		parseError(err),
	}
}

func GetSubjects(pool *pgx.ConnPool) DbResult {
	rows, err := pool.Query(GET_SUBJECTS)
	status := parseError(err)
	if status.IsError() {
		return DbResult{
			DbData{nil},
			status,
		}
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

	status = parseError(err)
	if status.IsError() {
		return DbResult{
			DbData{nil},
			status,
		}
	}

	return DbResult{
		DbData{subjects},
		status,
	}
}

func GetTgasBySubect(subject string, pool *pgx.ConnPool) DbResult {
	data, err := getTags(GET_TAGS_BY_SUBJECT, pool, subject)
	return DbResult{
		DbData{data},
		parseError(err),
	}
}
