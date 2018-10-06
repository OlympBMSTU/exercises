package main

import (
	"fmt"
	"net/http"

	"github.com/HustonMmmavr/excercieses/config"

	"github.com/HustonMmmavr/excercieses/controllers"

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

	pg_conf, err := config.GetPgConfigInstance()
	if err != nil {
		fmt.Println("eroor read pg_conf file")
		return nil, err
	}

	connPoolConfig := pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     pg_conf.GetHost(),
			User:     pg_conf.GetUser(),
			Password: pg_conf.GetPassword(),
			Database: pg_conf.GetDatabase(),
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
	http.HandleFunc("/api/excercieses/get", controllers.GetExcerciese(pool))

}

func main() {
	pool, err := Init()
	if err != nil {
		fmt.Println(err)
		panic("p")
	}

	InitRouter(pool)

	http.ListenAndServe("localhost:5469", nil)
}
