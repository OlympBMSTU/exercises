package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/OlympBMSTU/exercises/db"
	"github.com/OlympBMSTU/exercises/fstorage"
	"github.com/OlympBMSTU/exercises/parser"
	"github.com/OlympBMSTU/exercises/result"
	"github.com/OlympBMSTU/exercises/sender"
	"github.com/OlympBMSTU/exercises/views"
)

func UploadExerciseHandler(writer http.ResponseWriter, request *http.Request) {
	if !CheckMethodAndAuthenticate(writer, request, "POST") {
		return
	}

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

	dbEx := exView.ToExEntity()
	dbRes := db.SaveExercise(dbEx, request.Context())
	if dbRes.IsError() {
		WriteResponse(&writer, dbRes)
		// TODO delete file
		return
	}

	exID := uint(dbRes.GetData().(int))
	senderRes := sender.SendAnswer(exID, exView.Answer)
	if senderRes.IsError() {
		dbDelRes := db.DeleteExcerciese(exID, request.Context())
		fmt.Print(dbDelRes)
		//fsDelRes = fstorage.DeleteFile(filename)
		WriteResponse(&writer, senderRes)
		return
	}

	WriteResponse(&writer, dbRes)
}

func GetExercise(writer http.ResponseWriter, request *http.Request) {
	if !CheckMethodAndAuthenticate(writer, request, "GET") {
		return
	}

	// Get path variable from path
	idStr := strings.TrimPrefix(request.URL.Path, "/api/exercises/get/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(writer, "Incorrect path variable", http.StatusBadRequest)
		return
	}
	exID := uint(id)

	dbRes := db.GetExercise(exID, request.Context())
	WriteResponse(&writer, dbRes)
}

// GetExercises :
// This controller searches exercises in database by sended params:
// Path variables like /api/.../1/2/3
// Where
// 		1: subject
// 		2: tag
// 		3: level
// Also query variables
// 		1: limit
// 		2: offset
// 		3: order
func GetExercises(writer http.ResponseWriter, request *http.Request) {
	if !CheckMethodAndAuthenticate(writer, request, "POST") {
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
	order := query.Get("order")
	isDesc := false
	if order != "" && order == "desc" {
		isDesc = true
	}

	dbRes := db.GetExerciseList(tag, subject, level, limit, offset, isDesc, request.Context())
	WriteResponse(&writer, dbRes)
}

func UpdateExercise(writer http.ResponseWriter, request *http.Request) {
	// todo update tags, update file for after
	// update answer
	// update is broken level subject
	if !CheckMethodAndAuthenticate(writer, request, "POST") {
		return
	}

}
