package db

import (
	"bytes"
	"context"
	"fmt"
	"log"

	"github.com/OlympBMSTU/exercises/db/result"
	"github.com/OlympBMSTU/exercises/entities"
	"github.com/lib/pq"
)

const (
	DEFAULT_LIMIT  = 20
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
	limit int, offset int, order_level bool, ctx context.Context) result.DbResult {

	db := getDb(ctx)
	if db == nil {
		return result.ErrorResult(result.DB_CONN_ERROR, "")
	}

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

	rows, err := db.Query(query.String(), args...)
	defer rows.Close()

	if err != nil {
		log.Println(err.Error())
		return result.ErrorResult(err)
	}

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
		return result.ErrorResult(err)
	}

	updateData := existingEntity.GetDataForUpdateEntity(exEntity)

	if len(updateData) == 0 {
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
			// return result.
		}
		// search subject else no subject
	}

	tx, err := db.Begin()
	if err != nil {
		return result.ErrorResult(err)
	}

	// here starts transaction

	if updateData["new_tags"] != nil {
		tx.Exec("")
		delete(updateData, "new_tags")
		delete(updateData, "deleted_tags")
	}

	// firstly work with tags

	var query bytes.Buffer
	args := make([]interface{}, 1)
	for k, v := range updateData {
		args = append(args, v)
		query.WriteString(fmt.Sprintf("%s=$%d ", k, len(args)))
	}

	// after another update
	args = append(args, exEntity.Id)
	fmt.Print(query.String())
	query.WriteString(fmt.Sprintf("WHERE id=$", len(args)))
	tx.Exec(query.String(), args)
	err = tx.Commit()

	if err != nil {
		return result.ErrorResult(err)
	}

	return result.OkResult(existingEntity)
}

// todo how to check all check is not empty and others and its ok
// it's simply to check maybe we need to create function that returns map as changed values
// for ex that copies
// // query := "UPDATE exercises SET "
// rows, err := db.Query(GET_EXERCISE_BY_ID, exEntity.GetId())
// defer rows.Close()

// if err != nil {
// 	return result.ErrorResult(err)
// }

// var existingEntity *entities.ExerciseEntity
// for rows.Next() {
// 	excerciese, err = scanExercise(rows)
// }

// if err != nil {
// 	return result.ErrorResult(err)
// }

// var query bytes.Buffer
// query.WriteString(UPDATE_EXCERCISE)
// var args []interface{}

// fmt.Print(query)
// AuthorId uint
// FileName string
// Tags     []string
// Level    int
// Subject  string
// IsBroken bool
// if len(exEntity.FileName) > 0 && exEntity.FileName != existingEntity.FileName {
// 	args = append(args, exEntity.FileName)
// 	query.WriteString(fmt.Sprintf("filename=%d ", len(args))
// }

// if len(exEntity.Tags) > 0 {
// 	for i, tag := range exEntity.Tags {
// 		if tag != existingEntity.Tags[i] {
// 			args = append(args, exEntity.Tags)
// 			query.WriteString(fmt.Sprintf("tags=%d ", len(args))
// 			break
// 		}
// 	}
// }

// if exEntity.Level > 0 && exEntity.Level != existingEntity.Level {

// }n
