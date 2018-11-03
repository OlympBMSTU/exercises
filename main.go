package main

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/OlympBMSTU/exercises/config"
	"github.com/OlympBMSTU/exercises/controllers"
	"github.com/jackc/pgx"
)

type ContextInjector struct {
	ctx context.Context
	h   http.Handler
}

func (ci *ContextInjector) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	ci.h.ServeHTTP(writer, request.WithContext(ci.ctx))
}

func Init() (*pgx.ConnPool, error) {
	conf, err := config.GetConfigInstance()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	port, err := strconv.Atoi(conf.GetDBPort())
	if err != nil {
		log.Print(err)
		return nil, err
	}
	connPoolConfig := pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     conf.GetDBHost(),
			User:     conf.GetDBUser(),
			Port:     uint16(port),
			Password: conf.GetDBPassword(),
			Database: conf.GetDatabase(),
		},
		MaxConnections: 5,
	}

	pool, err := pgx.NewConnPool(connPoolConfig)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return pool, nil
}

func InitRouter(ctx context.Context) {
	http.Handle("/api/exercises/upload_exercise", &ContextInjector{ctx, http.HandlerFunc(controllers.UploadExerciseHandler)})
	http.Handle("/api/exercises/get/", &ContextInjector{ctx, http.HandlerFunc(controllers.GetExercise)})
	http.Handle("/api/exercises/list/", &ContextInjector{ctx, http.HandlerFunc(controllers.GetExercises)})
	http.Handle("/api/exercises/subjects/", &ContextInjector{ctx, http.HandlerFunc(controllers.GetSubjects)})
	http.Handle("/api/exercises/tags/", &ContextInjector{ctx, http.HandlerFunc(controllers.GetTags)})
}

func main() {
	pool, err := Init()
	if err != nil {
		log.Println(err.Error())
		panic("Error start server")
	}

	ctx := context.WithValue(context.Background(), "db", pool)

	InitRouter(ctx)

	http.ListenAndServe("localhost:5469", nil)
}
