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
	WriteResponse(&writer, "JSON", dbRes)
}

// firstly:
// maybe id in view set
// check new CheckMethodAndAuthenti

// parse form as one?  for create for update
// parse res new state that is not error, but empty maybe new method for view such as is empty
// good variant id sended in view if no id -> error
// if id exist, but no data ok then create file if no file and view is empty return error
// but which error return ? as parse error, as new error as CreateResponse
// maybe previous check that file exist and then parse form

// now db
// main prolem tags it could have
// 1) send tags all old
// 2) send tags all one or two changed its good variant
// 3) if sended one tag then its need to create one ol one new
// 4) isFalse its good
// now check all data fields
// thats not empty to db -> create query, fill array
// when update if such tag doesnt exist then delete it and tag exercise id
// else nothing to dospacspacesspaceses

// also todo for generating
// for subject
// for tag by class
// for firs level
// secons ... upto 10
// generate variant
// db stores filepath id and tags maybe if its need
// as array of part exercise
// fields are filled by old values if its changed then send

// frontend - list exercises
// list one
// also in list one you can change it
func UpdateExerciseHandler(writer http.ResponseWriter, request *http.Request) {
	// todo update tags, update file for after
	// update answer
	// update is broken level subject
	if !CheckMethodAndAuthenticate(writer, request, "POST") {
		return
	}

	var err error
	if err = request.ParseMultipartForm(-1); err != nil {
		WriteResponse(&writer, "JSON", map[string]string{
			"Message": "Error parse form",
			"Status":  "Error",
			"Data":    "",
		}, http.StatusBadRequest)
		// http.Error(writer, "Incorrect body", http.StatusBadRequest)
		return
	}

	WriteResponse(&writer, "JSON", map[string]interface{}{
		"Message": "Error parse form",
		"Status":  "Error",
		"Data":    nil,
	}, http.StatusBadRequest)
	return
	// parseRes := parser.ParseExViewPostUpdateForm(request.Form)
	// if parseRes.IsError() {
	// 	WriteResponse(&writer, "JSON", parseRes)
	// 	return
	// }

	// // exView := parseRes.GetData().(views.ExerciseView)
	// //
	// _, header, _ := request.FormFile("file")
	// // newFileName := ""

	// if header != nil {
	// 	fsRes := fstorage.WriteFile(header)
	// 	if fsRes.IsError() {
	// 		WriteResponse(&writer, "JSON", fsRes)
	// 		return
	// 	}
	// 	newFileName = fsRes.GetData().(string)
	// }

	// if exView.IsEmpty() && header == nil {

	// }

	// exView.SetFileName(fsRes.GetData().(string))

	//dbRes := db.UpdateExercise(id, exView)
	// if dbRes.IsError() {
	//
	// return
	//}

	// if answer changed sedn to

	// if send error

	// return

}
