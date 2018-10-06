package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/HustonMmmavr/excercieses/db"

	"github.com/HustonMmmavr/excercieses/entities"

	"github.com/HustonMmmavr/excercieses/config"

	"github.com/HustonMmmavr/excercieses/fstorage"
	"github.com/jackc/pgx"
)

type ExcercieseUpload struct {
	Answer     string `json:"answer"`
	FileBase64 string `json:"file"`
	FileName   string `json:"file_name"`
	author     uint64
}

func uploadExcercieseHandler(poll *pgx.ConnPool) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, requset *http.Request) {

		if requset.Method != "POST" {
			http.Error(writer, "Unsupported method", 404)
			return
		}

		if requset.Body == nil {
			http.Error(writer, "Please send a request body", 400)
			return
		}

		fmt.Println(poll)

		body, err := ioutil.ReadAll(requset.Body)
		defer requset.Body.Close()

		if err != nil {
			http.Error(writer, "Please send a request body", 400)
			return
		}

		var excerciese ExcercieseUpload
		err = json.Unmarshal(body, &excerciese)

		if err != nil {
			http.Error(writer, "Error json", 400)
			return
		}

		file, err := base64.StdEncoding.DecodeString(excerciese.FileBase64)
		if err != nil {
			http.Error(writer, "Incorrect file", 400)
			return
		}

		// get file name as fstorage
		newName := fstorage.ComputeName(excerciese.FileName)

		// write excerciese to db
		// check

		err = fstorage.WriteFile(file, newName, ".pdf")

		// excEnt := entities.NewExcercieseEntity(1, newName, "5")

		if err != nil {
			fmt.Println(err)
			return
		}

		if err != nil {
			http.Error(writer, "Error save file", 500)
			return
		}

		writer.Write([]byte("SUCCESS"))
	})
}

func main() {
	// http.HandleFunc("/api/excercieses/")
	conf, err := config.GetConfigInstance()
	fmt.Println(conf.GetFileStorageName())
	if err != nil {
		fmt.Println("Error in config file")
		return
	}

	pg_conf, err := config.GetPgConfigInstance()
	if err != nil {
		fmt.Println("eroor read pg_conf file")
	}

	connPoolConfig := pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     pg_conf.GetHost(),
			User:     pg_conf.GetUser(),
			Password: pg_conf.GetPassword(),
			Database: pg_conf.GetDatabase(),
		},
		MaxConnections: 5,
		// AfterConnect:   afterConnect,
	}

	pool, err := pgx.NewConnPool(connPoolConfig)
	if err != nil {
		fmt.Println("Error create pool")
		os.Exit(1)
	}

	http.HandleFunc("/api/excercieses/upload_excerciese", uploadExcercieseHandler(pool))

	fmt.Println(pool)

	existingFile, _ := os.Open("/home/mavr/Downloads/these.pdf")

	name := fstorage.ComputeName("these.pdf")
	bytes, _ := ioutil.ReadAll(existingFile)

	ex := entities.NewExcercieseEntity(1, name, "5")

	db.SaveExcerciese(ex, pool)
	fstorage.WriteFile(bytes, name, ".pdf")
}

// db cofing works
// 1) Create table
// 2) Create entity
// 3) save entity
// 4) get entity
// 5) get entities limit, offset
// 6) entity - предмет, тип, сложность
