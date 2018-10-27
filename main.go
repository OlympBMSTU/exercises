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
	// http.HandleFunc("/api/excercieses/get/", controllers.GetExcerciese(pool))
	// http.HandleFunc("/api/excercieses/list/", controllers.GetExcercieses(pool))
	http.HandleFunc("/api/excercieses/subjects/", controllers.GetSubjects(pool))
	// http.HandleFunc("/api/excercieses/tags/", controllers.GetTags(pool))
}

func main() {
	pool, err := Init()
	if err != nil {
		fmt.Println(err)
		panic("Error start server")
	}

	InitRouter(pool)

	// type ExcercieseEntity struct {
	// 	Id          uint
	// 	AuthorId    uint
	// 	FileName    string
	// 	RightAnswer string
	// 	Tags        []string
	// 	Level       uint
	// 	Subject     string
	// }

	// dt := [1,2]

	// ex := entities.ExcercieseEntity{
	// 	0,
	// 	1,
	// 	"sdfd",
	// 	"dsf",
	// 	[]string{"array", "data"},
	// 	3,
	// 	"mathematic",
	// }

	// res := db.SaveExcerciese(ex, pool)

	// res := db.GetExcerciese(1, pool)

	// fmt.Print(res)

	http.ListenAndServe("localhost:5469", nil)
}

// tests
// test create subjects correct

// create 3 subjects
// create ex ok
// create ex no subject
// drop table
// drop db
// drop connect

// test get tags by ex

// test get one ex ok
//   			not exist
//	maybe check tags for correct

// get excercises

// init db -> create
//

// ex := entities.ExcercieseEntity{
// 	0,
// 	1,
// 	"sdfd",
// 	"dsf",
// 	[]string{"array", "data"},
// 	3,
// 	"mathematics",
// }

// res := db.SaveExcerciese(ex, pool)
