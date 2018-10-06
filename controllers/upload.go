package controllers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/HustonMmmavr/excercieses/db"

	"github.com/HustonMmmavr/excercieses/fstorage"
	"github.com/HustonMmmavr/excercieses/views"
	"github.com/jackc/pgx"
)

func UploadExcercieseHandler(pool *pgx.ConnPool) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, requset *http.Request) {

		// url path
		fmt.Println(requset.URL.RequestURI)
		if requset.Method != "POST" {
			http.Error(writer, "Unsupported method", 404)
			return
		}

		if requset.Body == nil {
			http.Error(writer, "Please send a request body", 400)
			return
		}

		body, err := ioutil.ReadAll(requset.Body)
		defer requset.Body.Close()

		if err != nil {
			http.Error(writer, "Please send a request body", 400)
			return
		}

		var excerciese views.ExcercieseView
		err = json.Unmarshal(body, &excerciese)

		if err != nil {
			http.Error(writer, "Error json", 400)
			return
		}

		excercieseEntity := excerciese.ToEntity()

		file, err := base64.StdEncoding.DecodeString(excerciese.FileBase64)
		if err != nil {
			http.Error(writer, "Incorrect file", 400)
			return
		}

		// get file name as fstorage
		newName := fstorage.ComputeName(excerciese.FileName)

		err = fstorage.WriteFile(file, newName, ".pdf")
		excercieseEntity.SetFileName(newName)

		/////// ----- //////
		excercieseEntity.SetAuthor(0)

		db.SaveExcerciese(excercieseEntity, pool)

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

func GetExcerciese(pool *pgx.ConnPool) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		excerciese, err := db.GetExcerciese(14, pool)
		writer.Header().Set("Content-Type", "application/json")
		val, err := json.Marshal(excerciese)
		fmt.Println(err)
		writer.Write([]byte(val))
	})
}
