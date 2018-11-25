package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/OlympBMSTU/exercises/config"
	"github.com/OlympBMSTU/exercises/db"
	"github.com/OlympBMSTU/exercises/fstorage"
	"github.com/OlympBMSTU/exercises/parser"
	"github.com/OlympBMSTU/exercises/result"
	"github.com/OlympBMSTU/exercises/sender"
	"github.com/OlympBMSTU/exercises/views"
)

// UploadExerciseHandler : Controller that takes multipart form data
// parses it, saves exercise to db and sends answer to secret system
func UploadExerciseHandler(writer http.ResponseWriter, request *http.Request) {
	userID := CheckMethodAndAuthenticate(writer, request, "POST")
	if userID == nil {
		return
	}

	var err error
	if err = request.ParseMultipartForm(-1); err != nil {
		WriteResponse(&writer, "JSON", map[string]interface{}{
			"Message": "Error parse form",
			"Status":  "Error",
			"Data":    nil,
		}, http.StatusBadRequest)
		return
	}

	parseRes := parser.ParseExViewPostForm(request.Form)
	if parseRes.IsError() {
		WriteResponse(&writer, "JSON", parseRes)
		return
	}

	exView := parseRes.GetData().(views.ExerciseView)

	var fsRes result.Result
	_, header, _ := request.FormFile("file")

	fsRes = fstorage.WriteFile(header)
	if fsRes.IsError() {
		WriteResponse(&writer, "JSON", fsRes)
		return
	}

	exView.SetAuthor(*userID)
	exView.SetFileName(fsRes.GetData().(string))
	dbEx := exView.ToExEntity()
	dbRes := db.SaveExercise(dbEx, request.Context())
	if dbRes.IsError() {
		WriteResponse(&writer, "JSON", dbRes)
		// TODO delete file
		return
	}

	exID := uint(dbRes.GetData().(int))
	conf, _ := config.GetConfigInstance()
	if !conf.IsTest() {
		senderRes := sender.SendAnswer(exID, exView.Answer)
		if senderRes.IsError() {
			dbDelRes := db.DeleteExcerciese(exID, request.Context())
			fmt.Print(dbDelRes)
			//fsDelRes = fstorage.DeleteFile(filename)
			WriteResponse(&writer, "JSON", senderRes)
			return
		}
	}

	WriteResponse(&writer, "JSON", dbRes)
}

// GetExerciseHandler : controller that searches exercise in database
// by presented exercise id as path variable (ex /api/../id)
func GetExerciseHandler(writer http.ResponseWriter, request *http.Request) {
	userID := CheckMethodAndAuthenticate(writer, request, "GET")
	if userID == nil {
		return
	}

	// Get path variable from path
	idStr := strings.TrimPrefix(request.URL.Path, "/api/exercises/get/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		WriteResponse(&writer, "JSON", map[string]interface{}{
			"Message": "Incorrect path variable",
			"Status":  "Error",
			"Data":    nil,
		}, http.StatusBadRequest)
		return
	}
	exID := uint(id)

	dbRes := db.GetExercise(exID, request.Context())
	WriteResponse(&writer, "JSON", dbRes)
}

// GetExercises : controller that searches exercises in database by presented conditio params:
// Path variables like /api/.../subject/tag/level
// Where
// 		1: subject - string
// 		2: tag 	   - string
// 		3: level   - integer
// Also query variables
// 		1: limit   - integer
// 		2: offset  - integer
// 		3: order   - string
func GetExercises(writer http.ResponseWriter, request *http.Request) {
	userID := CheckMethodAndAuthenticate(writer, request, "GET")
	if userID == nil {
		return
	}

	query := request.URL.Query()
	pathVariablesStr := strings.TrimPrefix(request.URL.Path, "/api/exercises/list/")
	vars := strings.Split(pathVariablesStr, "/")
	subject := ""
	tag := ""
	level := -1

	if len(vars) == 0 || (len(vars) == 1 && vars[0] == "") {
		WriteResponse(&writer, "JSON", map[string]interface{}{
			"Message": "Not enough parameters for request",
			"Status":  "Error",
			"Data":    nil,
		}, http.StatusBadRequest)
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
				WriteResponse(&writer, "JSON", map[string]interface{}{
					"Message": "Incorrect level variable",
					"Status":  "Error",
					"Data":    nil,
				}, http.StatusBadRequest)
				return
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

	isBrokenStr := query.Get("is_broken")
	isBroken := false
	if isBrokenStr != "" && isBrokenStr == "true" {
		isBroken = true
	}

	dbRes := db.GetExerciseList(tag, subject, level, limit, offset, isDesc, isBroken, request.Context())
	WriteResponse(&writer, "JSON", dbRes)
}

func UpdateExerciseHandler(writer http.ResponseWriter, request *http.Request) {
	userID := CheckMethodAndAuthenticate(writer, request, "POST")
	if userID == nil {
		return
	}

	var err error
	if err = request.ParseMultipartForm(-1); err != nil {
		WriteResponse(&writer, "JSON", map[string]interface{}{
			"Message": "Error parse form",
			"Status":  "Error",
			"Data":    nil,
		}, http.StatusBadRequest)
		return
	}

	parseRes := parser.ParseExViewPostUpdateForm(request.Form)
	if parseRes.IsError() {
		WriteResponse(&writer, "JSON", parseRes)
		return
	}

	// about id int or uint
	exView := parseRes.GetData().(views.ExerciseView)

	_, header, _ := request.FormFile("file")
	if header != nil {
		fsRes := fstorage.WriteFile(header)
		if fsRes.IsError() {
			WriteResponse(&writer, "JSON", fsRes)
			return
		}
		exView.SetFileName(fsRes.GetData().(string))
	}

	// if old file == new file names doesnt mathc, so if update we need to send new file else no
	entity := exView.ToExEntity()

	dbRes := db.UpdateExercise(entity, request.Context())

	// OR create additional error NOT_UPDATED and map
	if dbRes.GetData() == nil && !dbRes.IsError() {
		WriteResponse(&writer, "JSON",
			map[string]interface{}{
				"Message": "No new data added to exercise",
				"Data":    nil,
				"Status":  "Error",
			}, http.StatusBadRequest)
		return
	}
	// also here smtp
	WriteResponse(&writer, "JSON", dbRes)
}
