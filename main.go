package main

import (
	"fmt"
	"net/http"

	"github.com/OlympBMSTU/exercises/config"
	"github.com/OlympBMSTU/exercises/controllers"
	"github.com/jackc/pgx"
)

func Init() (*pgx.ConnPool, error) {
	conf, err := config.GetConfigInstance()
	if err != nil {
		fmt.Println("Error in config file")
		return nil, err
	}

	connPoolConfig := pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     conf.GetDBHost(),
			User:     conf.GetDBUser(),
			Password: conf.GetDBPassword(),
			Database: conf.GetDatabase(),
		},
		MaxConnections: 5,
	}

	pool, err := pgx.NewConnPool(connPoolConfig)
	if err != nil {
		fmt.Println("Error create pool")
		return nil, err
	}
	return pool, nil
}

func InitRouter(pool *pgx.ConnPool) {
	http.HandleFunc("/api/exercises/upload_exercise", controllers.UploadExerciseHandler(pool))
	http.HandleFunc("/api/exercises/get/", controllers.GetExercise(pool))
	http.HandleFunc("/api/exercises/list/", controllers.GetExercises(pool))
	http.HandleFunc("/api/exercises/subjects/", controllers.GetSubjects(pool))
	http.HandleFunc("/api/exercises/tags/", controllers.GetTags(pool))
}

func main() {
	pool, err := Init()
	if err != nil {
		fmt.Println(err)
		panic("Error start server")
	}

	InitRouter(pool)

	http.ListenAndServe("localhost:5469", nil)
}
