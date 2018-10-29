package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/OlympBMSTU/exercises/parser"
	"github.com/OlympBMSTU/exercises/sender"

	"github.com/OlympBMSTU/exercises/auth"
	"github.com/OlympBMSTU/exercises/db"
	"github.com/OlympBMSTU/exercises/fstorage"
	"github.com/OlympBMSTU/exercises/result"
	"github.com/OlympBMSTU/exercises/views"
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

		parseRes := parser.ParseExViewPostForm(request.Form)
		if parseRes.IsError() {
			WriteResponse(&writer, parseRes)
			return
		}

		exView := parseRes.GetData().(views.ExerciseView)

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

		exView.SetFileName(fsRes.GetData().(string))
		// exView.SetAuthor(authRes.GetData().(uint))
		exView.SetAuthor(0)

		dbEx := exView.ToExEntity()
		dbRes := db.SaveExercise(dbEx, pool)
		if dbRes.IsError() {
			WriteResponse(&writer, dbRes)
			// TODO delete file
			return
		}

		exId := uint(dbRes.GetData().(int))
		senderRes := sender.SendAnswer(exId, exView.Answer)
		if senderRes.IsError() {
			//dbDelRes = db.DeleteExcerciese(exId, pool)
			//fsDelRes = fstorage.DeleteFile(filename)
			WriteResponse(&writer, senderRes)
			return
		}

		WriteResponse(&writer, dbRes)
	})
}

func GetExercise(pool *pgx.ConnPool) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		OptionsCredentials(&writer)
		if request.Method == "OPTIONS" {
			writer.Write([]byte("hi"))
			return
		}

		if request.Method != "GET" {
			http.Error(writer, "Unsopported method", 405)
			return
		}

		writer.Header().Set("Content-Type", "application/json")

		authRes := auth.AuthByUserCookie(request, "bmstuOlympAuth")
		if authRes.IsError() {
			WriteResponse(&writer, authRes)
			return
		}

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
		OptionsCredentials(&writer)
		if request.Method == "OPTIONS" {
			writer.Write([]byte("hi"))
			return
		}

		if request.Method != "GET" {
			http.Error(writer, "Unsopported method", 405)
			return
		}

		writer.Header().Set("Content-Type", "application/json")

		authRes := auth.AuthByUserCookie(request, "bmstuOlympAuth")
		if authRes.IsError() {
			WriteResponse(&writer, authRes)
			return
		}

		query := request.URL.Query()
		pathVariablesStr := strings.TrimPrefix(request.URL.Path, "/api/exercises/list/")
		vars := strings.Split(pathVariablesStr, "/")
		subject := ""
		tag := ""
		level := -1

		if len(vars) == 0 || (len(vars) == 1 && vars[0] == "") {
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
