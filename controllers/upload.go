package controllers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/OlympBMSTU/excericieses/db"
	"github.com/OlympBMSTU/excericieses/fstorage"
	"github.com/OlympBMSTU/excericieses/views"
	"github.com/jackc/pgx"
)

func UploadExcercieseHandler(pool *pgx.ConnPool) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		if request.Method != "POST" {
			http.Error(writer, "Unsupported method", 404)
			return
		}

		if request.Body == nil {
			http.Error(writer, "Please send a request body", 400)
			return
		}

		// also here we got author id
		// if !AuthUser(cookie) {
		// 	http.Error(writer, "Please authorize", 403)
		// 	return
		// }

		body, err := ioutil.ReadAll(request.Body)
		defer request.Body.Close()

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

		// represent name in file storage
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

		if request.Method != "GET" {
			http.Error(writer, "Unsopported method", 404)
			return
		}

		// Get path variable from path
		idStr := strings.TrimPrefix(request.URL.Path, "/api/excercieses/get/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(writer, "Incorrect path variable", 404)
		}

		uId := uint(id)

		excerciese, err := db.GetExcerciese(uId, pool)
		writer.Header().Set("Content-Type", "application/json")
		val, err := json.Marshal(excerciese)
		fmt.Println(err)
		writer.Write([]byte(val))
	})
}
