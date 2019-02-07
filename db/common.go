package db

import (
	"context"
	"log"
	"time"

	"github.com/OlympBMSTU/exercises/entities"
	"github.com/OlympBMSTU/exercises/logger"
	"github.com/jackc/pgx"
)

func getDb(ctx context.Context) *pgx.ConnPool {
	switch ctx.Value("db").(type) {
	case *pgx.ConnPool:
		return ctx.Value("db").(*pgx.ConnPool)
	default:
		log.Println("Db not set in context")
		return nil
	}
}

func scanExercise(rows *pgx.Rows) (*entities.ExerciseEntity, error) {
	log := logger.GetLogger()
	var exercise entities.ExerciseEntity
	var answers []byte
	var data time.Time
	err := rows.Scan(
		&exercise.Id,
		&exercise.AuthorId,
		&exercise.Level,
		&exercise.FileName,
		&exercise.Subject,
		&exercise.Tags,
		&exercise.IsBroken,
		&exercise.Class,
		&exercise.Position,
		&exercise.Mark,
		&exercise.TypeOlymp,
		&answers,
		&data,
	)
	exercise.Created = data.String()

	if err != nil {
		// fmt.Print(err.Error())
		log.Error("Error parse exercise", err)
		return nil, err
	}

	return &exercise, nil
}

func getTags(query string, pool *pgx.ConnPool, args ...interface{}) (*[]string, error) {
	rows, err := pool.Query(query, args[0])
	defer rows.Close()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var tags []string

	for rows.Next() {
		var tag string
		err := rows.Scan(&tag)
		if err != nil {
			log.Println(err.Error())
			continue
		}
		tags = append(tags, tag)
	}

	return &tags, nil
}
