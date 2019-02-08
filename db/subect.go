package db

import (
	"context"

	"github.com/OlympBMSTU/exercises/db/result"
	"github.com/OlympBMSTU/exercises/logger"
)

func AddSubject(subject string, ctx context.Context) result.DbResult {
	log := logger.GetLogger()
	db := getDb(ctx)
	if db == nil {
		log.Error("No db finded", nil)
		return result.ErrorResult(result.DB_CONN_ERROR, "")
	}
	_, err := db.Exec(ADD_SUBJECT, subject)
	if err != nil {
		log.Error("Error add subject", err)
	}
	return result.CreateResult(nil, err, result.CREATED)
}

func GetSubjects(ctx context.Context) result.DbResult {
	log := logger.GetLogger()
	db := getDb(ctx)
	if db == nil {
		log.Error("No db finded", nil)
		return result.ErrorResult(result.DB_CONN_ERROR, "")
	}

	rows, err := db.Query(GET_SUBJECTS)
	if err != nil {
		log.Error("Error get subjects", err)
		return result.ErrorResult(err)
	}

	var subjects []string
	for rows.Next() {
		var subject string

		err = rows.Scan(&subject)
		if err != nil {
			log.Error("Error scan subject", err)
			break
		}

		subjects = append(subjects, subject)
	}

	// wtf
	if err != nil {
		log.Error("Error scan subject", err)
		return result.ErrorResult(err)
	}

	if len(subjects) == 0 {
		log.Warn("Empty result")
		return result.ErrorResult(result.EMPTY_RESULT, "")
	}

	return result.OkResult(subjects)
}

func GetTgasBySubect(subject string, ctx context.Context) result.DbResult {
	log := logger.GetLogger()
	db := getDb(ctx)
	if db == nil {
		log.Error("Error get subjects", nil)
		return result.ErrorResult(result.DB_CONN_ERROR, "")
	}

	data, err := getTags(GET_TAGS_BY_SUBJECT, db, subject)
	if err != nil {
		log.Error("Error get tags", err)
		return result.ErrorResult(err)
	}

	if len(*data) == 0 {
		log.Warn("Empty result")
		return result.ErrorResult(result.EMPTY_RESULT, "")
	}

	return result.OkResult(data)
}

func SaveSubject(subject string, ctx context.Context) result.DbResult {
	log := logger.GetLogger()
	db := getDb(ctx)
	if db == nil {
		log.Error("Error get subjects", nil)
		return result.ErrorResult(result.DB_CONN_ERROR, "")
	}

	row := db.QueryRow(ADD_SUBJECT, subject)

	var subjID int
	err := row.Scan(&subjID)

	if err != nil {
		log.Error("Error save subjects", err)
		return result.ErrorResult(err)
	}

	return result.OkResult(subjID)
}
