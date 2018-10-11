package db

import "github.com/jackc/pgx"

func AddSubject(subject string, pool *pgx.ConnPool) error {
	_, err := pool.Exec(ADD_SUBJECT, subject)
	return err
}

func GetSubjects(pool *pgx.ConnPool) (*[]string, error) {
	rows, err := pool.Query(GET_SUBJECTS)
	if err != nil {
		return nil, err
	}

	var subjects []string
	for rows.Next() {
		var subject string

		err := rows.Scan(&subject)
		if err != nil {
			return nil, err
		}

		subjects = append(subjects, subject)
	}

	return &subjects, nil
}
