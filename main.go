package main

import (
	"fmt"
	"net/http"

	"github.com/OlympBMSTU/excericieses/config"
	"github.com/OlympBMSTU/excericieses/controllers"
	"github.com/jackc/pgx"
)

type ExcercieseUpload struct {
	Answer     string `json:"answer"`
	FileBase64 string `json:"file"`
	FileName   string `json:"file_name"`
	author     uint64
}

func Init() (*pgx.ConnPool, error) {
	conf, err := config.GetConfigInstance()
	fmt.Println(conf.GetFileStorageName())
	if err != nil {
		fmt.Println("Error in config file")
		return nil, err
	}

	connPoolConfig := pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     conf.GetHost(),
			User:     conf.GetUser(),
			Password: conf.GetPassword(),
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
	http.HandleFunc("/api/excercieses/upload_excerciese", controllers.UploadExcercieseHandler(pool))
	http.HandleFunc("/api/excercieses/get/", controllers.GetExcerciese(pool))
	http.HandleFunc("/api/excercieses/list/", controllers.GetExcercieses(pool))
	http.HandleFunc("/api/excercieses/subjects/", controllers.GetSubjects(pool))
	http.HandleFunc("/api/excercieses/tags/", controllers.GetTags(pool))
	// http.HandleFunc("/api/excercieses/subjects/")
	// http.HandleFunc("/api/excercieses/") tag
	// count
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
