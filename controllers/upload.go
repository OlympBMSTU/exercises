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

// func getDataFromParameter(param []string) (bool, *string) {
// 	if len(param) > 0 &&
// }

func GetExcercieses(pool *pgx.ConnPool) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != "GET" {
			http.Error(writer, "Unsopported method", 404)
			return
		}
		query := request.URL.Query()
		pathVariablesStr := strings.TrimPrefix(request.URL.Path, "/api/excercieses/list/")
		vars := strings.Split(pathVariablesStr, "/")
		subject := ""
		tag := ""
		level := -1

		if len(vars) == 0 {
			http.Error(writer, "Not enough parameter", 404)
			return
		}

		// its very scary !!!!!
		for i, data := range vars {
			if i == 0 {
				subject = vars[0]
			}
			if i == 1 {
				tag = vars[1]
			}
			if i == 2 && data != "" {
				var err error
				level, err = strconv.Atoi(data)
				if err != nil {
					http.Error(writer, "INCORRECT PATH", 404)
				}

			}
		}

		limitArr := query["limit"]
		limit := -1
		if len(limitArr) > 0 {
			limit, _ = strconv.Atoi(limitArr[0])
		}
		offsetArr := query["offset"]
		offset := -1
		if len(offsetArr) > 0 {
			offset, _ = strconv.Atoi(limitArr[0])
		}

		fmt.Println(limit, offset)

		// its fucking crutch maybe, todo refactor !!!!!!!!!

		// check order for quer
		order := query["order"]
		is_desc := false
		if len(order) > 0 && order[0] == "desc" {
			is_desc = true
		}

		// 1 - subject 2 - tag 3 - level
		// query 1 - limit 2 - offset 3 - order

		entities, err := db.GetExcercieseList(tag, subject, level, limit, offset, is_desc, pool)
		if err != nil {
			http.Error(writer, "Internal server err", 500)
			return
		}

		val, err := json.Marshal(entities)
		fmt.Println(err)
		writer.Write([]byte(val))

	})
}
