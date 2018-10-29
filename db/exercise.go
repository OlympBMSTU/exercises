package db

import (
	"bytes"
	"fmt"

	"github.com/OlympBMSTU/exercises/db/result"
	"github.com/OlympBMSTU/exercises/entities"
	"github.com/jackc/pgx"
	"github.com/lib/pq"
)

const (
	DEFAULT_LIMIT  = 20
	DEFAULT_OFFSET = 0
)

func DeleteExcerciese(exId uint, pool *pgx.ConnPool) result.DbResult {
	_, err := pool.Exec(DELETE_EXERCISE, exId)
	return result.CreateResult(nil, err)
}

func SaveExercise(exercise entities.ExerciseEntity, pool *pgx.ConnPool) result.DbResult {
	row := pool.QueryRow(INSERT_EXERCISE,
		exercise.GetAuthorId(),
		exercise.GetLevel(),
		exercise.GetFileName(),
		exercise.GetSubject(),
		pq.Array(exercise.GetTags()),
	)

	var returnedId int
	err := row.Scan(&returnedId)

	if err != nil {
		return result.ErrorResult(err)
	}

	if returnedId == -1 {
		return result.ErrorResult(result.NO_SUBJECT_ERROR, "There is no subject in db")
	}

	return result.OkResult(returnedId, result.CREATED)
}

func GetExercise(id uint, pool *pgx.ConnPool) result.DbResult {
	rows, err := pool.Query(GET_EXERCISE_BY_ID, id)
	defer rows.Close()

	if err != nil {
		return result.ErrorResult(err)
	}

	var excerciese *entities.ExerciseEntity
	for rows.Next() {
		excerciese, err = scanExercise(rows)
	}

	if err != nil {
		return result.ErrorResult(err)
	}

	if excerciese == nil {
		return result.ErrorResult(result.EMPTY_RESULT, "")
	}

	return result.OkResult(entities.NewExRepresentation(*excerciese))
}

func GetExerciseList(tag string, subject string, level int,
	limit int, offset int, order_level bool, poll *pgx.ConnPool) result.DbResult {

	var query bytes.Buffer

	var args []interface{}
	if tag != "" {
		query.WriteString(GET_EXERCISE_BY_SUBJECT_AND_TAG)
		args = append(args, subject, tag, subject)
	} else {
		args = append(args, subject)
		query.WriteString(GET_EXERCISES_BY_SUBJECT)
	}

	if level != -1 {
		args = append(args, level)
		query.WriteString(fmt.Sprintf("AND ex.level = $%d ", len(args)))
	} else {
		query.WriteString("ORDER BY ex.level ")
		if order_level {
			query.WriteString("DESC ")
		}
	}

	if limit == -1 {
		limit = DEFAULT_LIMIT
	}

	args = append(args, limit)
	query.WriteString(fmt.Sprintf("LIMIT $%d ", len(args)))

	if offset == -1 {
		offset = DEFAULT_OFFSET
	}

	args = append(args, offset)
	query.WriteString(fmt.Sprintf("OFFSET $%d ", len(args)))

	rows, err := poll.Query(query.String(), args...)
	defer rows.Close()

	if err != nil {
		return result.ErrorResult(err)
	}

	// need to send tags with excercieses,
	// or front will do request for this
	var entities []entities.ExerciseEntity
	for rows.Next() {
		exercise, err := scanExercise(rows)

		if err != nil {
			return result.ErrorResult(err)
		}
		entities = append(entities, *exercise)
	}

	if len(entities) == 0 {
		return result.ErrorResult(result.EMPTY_RESULT, "")
	}

	return result.OkResult(entities)
}
