package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/OlympBMSTU/excericieses/auth"
	"github.com/OlympBMSTU/excericieses/db"
	"github.com/OlympBMSTU/excericieses/entities"
	"github.com/OlympBMSTU/excericieses/fstorage"
	"github.com/OlympBMSTU/excericieses/result"
	"github.com/jackc/pgx"
)

func UploadExerciseHandler(pool *pgx.ConnPool) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		OptionsCredentials(&writer)
		if request.Method == "OPTIONS" {
			writer.Write([]byte("hi"))
			return
		}

		if request.Method != "POST" {
			http.Error(writer, "Unsupported method", 405)
			return
		}

		authRes := auth.AuthByUserCookie(request, "bmstuOlympAuth")
		if authRes.IsError() {
			WriteResponse(&writer, authRes)
			return
		}
		writer.Header().Set("Content-Type", "application/json")

		var err error
		if err = request.ParseMultipartForm(-1); err != nil {
			http.Error(writer, "Incorrect body", http.StatusBadRequest)
			return
		}

		answer := request.Form["answer"]
		subject := request.Form["subject"]
		level_str := request.Form["level"]
		tags_arr := request.Form["tags"]

		if len(answer) < 1 || len(subject) < 1 || len(level_str) < 1 || len(tags_arr) < 1 {
			http.Error(writer, "Incorrect body", http.StatusBadRequest)
			return
		}

		var tags []string
		err = json.Unmarshal([]byte(tags_arr[0]), &tags)
		if err != nil {
			http.Error(writer, "Incorrect tags in body", http.StatusBadRequest)
			return
		}

		var level int
		if level, err = strconv.Atoi(level_str[0]); err != nil {
			http.Error(writer, "Incorrect body", http.StatusBadRequest)
			return
		}

		var fsRes result.Result
		for _, fheaders := range request.MultipartForm.File {
			for _, hdr := range fheaders {
				fsRes = fstorage.WriteFile(hdr)
				if fsRes.IsError() {
					WriteResponse(&writer, fsRes)
					return
				}
			}
		}

		filename := fsRes.GetData().(string)
		author_id := 0
		dbExcerciese := entities.NewExerciseEntity(uint(author_id), filename, answer[0],
			tags, uint(level), subject[0])

		// err := sender.SendAnswer(0, "hi")
		// if err := nil {
		// 	// db.RemoveExcerciese(excercieseEntity.Id)
		// 	// fs.RemoveFile(newName)
		// 	return
		// }

		dbRes := db.SaveExercise(dbExcerciese, pool)
		WriteResponse(&writer, dbRes)
	})
}

func GetExercise(pool *pgx.ConnPool) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		if request.Method != "GET" {
			http.Error(writer, "Unsopported method", 405)
			return
		}
		writer.Header().Set("Content-Type", "application/json")

		// Get path variable from path
		idStr := strings.TrimPrefix(request.URL.Path, "/api/exercises/get/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(writer, "Incorrect path variable", http.StatusBadRequest)
			return
		}
		exId := uint(id)

		dbRes := db.GetExercise(exId, pool)
		WriteResponse(&writer, dbRes)
	})
}

func GetExercises(pool *pgx.ConnPool) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != "GET" {
			http.Error(writer, "Unsopported method", 405)
			return
		}

		writer.Header().Set("Content-Type", "application/json")

		query := request.URL.Query()
		pathVariablesStr := strings.TrimPrefix(request.URL.Path, "/api/exercises/list/")
		vars := strings.Split(pathVariablesStr, "/")
		subject := ""
		tag := ""
		level := -1

		if len(vars) == 0 {
			http.Error(writer, "Not enough parameter", 404)
			return
		}

		// its very scary !!!!!  // use reflect and refactor level
		for i, data := range vars {
			if i == 0 {
				subject = vars[i]
			}
			if i == 1 {
				tag = vars[i]
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
			offset, _ = strconv.Atoi(offsetArr[0])
		}

		// its fucking crutch maybe, todo refactor !!!!!!!!!

		// check order for quer
		order := query["order"]
		is_desc := false
		if len(order) > 0 && order[0] == "desc" {
			is_desc = true
		}

		// 1 - subject 2 - tag 3 - level
		// query 1 - limit 2 - offset 3 - order

		dbRes := db.GetExerciseList(tag, subject, level, limit, offset, is_desc, pool)
		WriteResponse(&writer, dbRes)
	})
}
