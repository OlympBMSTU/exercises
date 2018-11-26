package db

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/OlympBMSTU/exercises/db/result"
	"github.com/OlympBMSTU/exercises/entities"
	"github.com/lib/pq"
)

const (
	DEFAULT_LIMIT  = 300
	DEFAULT_OFFSET = 0
)

func DeleteExcerciese(exId uint, ctx context.Context) result.DbResult {
	db := getDb(ctx)
	if db == nil {
		return result.ErrorResult(result.DB_CONN_ERROR, "")
	}
	_, err := db.Exec(DELETE_EXERCISE, exId)
	if err != nil {
		log.Println(err.Error())
	}
	return result.CreateResult(nil, err)
}

func SaveExercise(exercise entities.ExerciseEntity, ctx context.Context) result.DbResult {
	db := getDb(ctx)
	if db == nil {
		return result.ErrorResult(result.DB_CONN_ERROR, "")
	}

	row := db.QueryRow(INSERT_EXERCISE,
		exercise.GetAuthorId(),
		exercise.GetLevel(),
		exercise.GetFileName(),
		exercise.GetSubject(),
		pq.Array(exercise.GetTags()),
	)

	var returnedId int
	err := row.Scan(&returnedId)

	if err != nil {
		log.Println(err.Error())
		return result.ErrorResult(err)
	}

	if returnedId == -1 {
		log.Println("No subject in db")
		return result.ErrorResult(result.NO_SUBJECT_ERROR, "There is no subject in db")
	}

	return result.OkResult(returnedId, result.CREATED)
}

func GetExercise(id uint, ctx context.Context) result.DbResult {
	db := getDb(ctx)
	if db == nil {
		return result.ErrorResult(result.DB_CONN_ERROR, "")
	}

	rows, err := db.Query(GET_EXERCISE_BY_ID, id)
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
		log.Println("Result set is empty")
		return result.ErrorResult(result.EMPTY_RESULT, "")
	}

	return result.OkResult(entities.NewExRepresentation(*excerciese))
}

func GetExerciseList(tag string, subject string, level int,
	limit int, offset int, order_level bool, isBroken bool, ctx context.Context) result.DbResult {

	db := getDb(ctx)
	if db == nil {
		return result.ErrorResult(result.DB_CONN_ERROR, "")
	}

	// var query bytes.Buffer
	query := ""

	var args []interface{}
	if tag != "" {
		query = GET_EXERCISE_BY_SUBJECT_AND_TAG
		// query.WriteString(GET_EXERCISE_BY_SUBJECT_AND_TAG)
		args = append(args, subject, tag, subject)
	} else {
		args = append(args, subject)
		query = GET_EXERCISES_BY_SUBJECT
		// query.WriteString(GET_EXERCISES_BY_SUBJECT)
	}

	// args = append(args, isBroken)
	// query += fmt.Sprintf("AND ex.is_broken =$%d ", len(args))
	if level != -1 {
		args = append(args, level)
		query += fmt.Sprintf("AND ex.level = $%d ", len(args))
		// query.WriteString(fmt.Sprintf("AND ex.level = $%d ", len(args)))
	} else {
		query += "ORDER BY ex.level, ex.id "

		// query.WriteString("ORDER BY ex.level ")
		if order_level {
			query += "DESC "
			//query.WriteString("DESC ")
		}
	}

	if limit == -1 {
		limit = DEFAULT_LIMIT
	}

	args = append(args, limit)
	query += fmt.Sprintf("LIMIT $%d ", len(args))
	// query.WriteString(fmt.Sprintf("LIMIT $%d ", len(args)))

	if offset == -1 {
		offset = DEFAULT_OFFSET
	}

	args = append(args, offset)
	query += fmt.Sprintf("OFFSET $%d ", len(args))
	// query.WriteString(fmt.Sprintf("OFFSET $%d ", len(args)))

	rows, err := db.Query(query, args...)
	defer rows.Close()

	if err != nil {
		log.Println(err.Error())
		return result.ErrorResult(err)
	}
	fmt.Print(query)
	// need to send tags with excercieses,
	// or front will do request for this
	var entities []entities.ExerciseEntity
	for rows.Next() {
		exercise, err := scanExercise(rows)

		if err != nil {
			log.Println(err.Error())
			return result.ErrorResult(err)
		}
		entities = append(entities, *exercise)
	}

	if len(entities) == 0 {
		return result.ErrorResult(result.EMPTY_RESULT, "")
	}

	return result.OkResult(entities)
}

func UpdateExercise(exEntity entities.ExerciseEntity, ctx context.Context) result.DbResult {
	db := getDb(ctx)
	if db == nil {
		log.Print("No db specified")
		return result.ErrorResult(result.DB_CONN_ERROR, "")
	}

	rows, err := db.Query(GET_EXERCISE_BY_ID, exEntity.GetId())
	defer rows.Close()

	if err != nil {
		return result.ErrorResult(err)
	}

	var existingEntity *entities.ExerciseEntity
	for rows.Next() {
		existingEntity, err = scanExercise(rows)
	}

	if err != nil {
		log.Println("cant read rows:", err)
		return result.ErrorResult(err)
	}

	if existingEntity == nil {
		log.Println("No such entity with id " + string(exEntity.Id))
		return result.ErrorResult(result.EMPTY_RESULT, "")
	}

	updateData := existingEntity.GetDataForUpdateEntity(exEntity)

	if len(updateData) == 0 {
		log.Print("Nothing to update")
		return result.OkResult(nil)
	}

	if updateData["subject"] != nil {
		row := db.QueryRow(GET_COUNT_SUBJECT, updateData["subject"])
		cnt := 0
		err = row.Scan(&cnt)
		if err != nil {
			return result.ErrorResult(err)
		}

		if cnt == 0 {
			log.Println("No subject in db")
			return result.ErrorResult(result.NO_SUBJECT_ERROR, "There is no subject in db")
		}
		// search subject else no subject
	}

	tx, err := db.Begin()
	if err != nil {
		log.Print("Cant create transaction: ", err)
		return result.ErrorResult(err)
	}

	// here starts transaction

	if updateData["tags_to_add"] != nil || updateData["tags_to_delete"] != nil {
		_, err = tx.Exec(UPDATE_TAGS_BY_EX, exEntity.Id, existingEntity.Subject, exEntity.Subject, updateData["tags_to_add"], updateData["tags_to_remove"])

		if err != nil {
			log.Print("cant update tags", err)
			return result.ErrorResult(err)
		}
		delete(updateData, "tags_to_add")
		delete(updateData, "tags_to_remove")
	}

	// firstly work with tags

	query := UPDAET_EXERCISE
	args := make([]interface{}, 0)
	for k, v := range updateData {
		args = append(args, v)
		query += fmt.Sprintf("%s=$%d,", k, len(args))
	}

	query = strings.TrimRight(query, ",")
	// after another update
	args = append(args, exEntity.Id)
	query += fmt.Sprintf(" WHERE id=$%d", len(args))
	_, err = tx.Exec(query, args...)
	// fmt.Print(query)

	if err != nil {
		log.Print("Query: ", query)
		log.Print("Error: ", err)
		return result.ErrorResult(err)
	}
	err = tx.Commit()

	if err != nil {
		log.Print("Error: ", err)
		return result.ErrorResult(err)
	}

	return result.OkResult(existingEntity)
}
