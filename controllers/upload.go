package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/OlympBMSTU/excericieses/auth"
	matcher "github.com/OlympBMSTU/excericieses/controllers/matcher_result"
	"github.com/OlympBMSTU/excericieses/db"
	"github.com/OlympBMSTU/excericieses/entities"
	"github.com/OlympBMSTU/excericieses/fstorage"
	"github.com/OlympBMSTU/excericieses/result"
	"github.com/jackc/pgx"
)

func m(d map[string][]string) {
	fmt.Print(d)
}

func UploadExcercieseHandler(pool *pgx.ConnPool) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		if request.Method != "POST" {
			http.Error(writer, "Unsupported method", 405)
			return
		}

		var err error

		cookie, err := request.Cookie("bmstuOlympAuth")
		// unauth
		if err != nil {
			http.Error(writer, "Unauthorized", 403)
			return
		}

		auth_res := auth.AuthUser(cookie.Value)
		if auth_res.IsError() {
			httpRes := matcher.MatchResult(auth_res)
			writer.WriteHeader(httpRes.GetStatus())
			writer.Write(httpRes.GetData())
			return
		}

		if err = request.ParseMultipartForm(-1); err != nil {
			http.Error(writer, "Incorrect body", http.StatusBadRequest)
			return
		}

		answer := request.Form["answer"]
		subject := request.Form["subject"]
		level_str := request.Form["level"]
		tags := request.Form["tags"]

		if len(answer) < 1 || len(subject) < 1 || len(level_str) < 1 || len(tags) < 1 {
			http.Error(writer, "Incorrect body", http.StatusBadRequest)
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
					http_res := matcher.MatchResult(fsRes)
					writer.WriteHeader(http_res.GetStatus())
					writer.Write(http_res.GetData())
					return
				}
			}
		}

		filename := fsRes.GetData().(string)
		author_id := 0
		dbExcerciese := entities.NewExcercieseEntity(uint(author_id), filename, answer[0],
			tags, uint(level), subject[0])

		// err := sender.SendAnswer(0, "hi")
		// if err := nil {
		// 	// db.RemoveExcerciese(excercieseEntity.Id)
		// 	// fs.RemoveFile(newName)
		// 	return
		// }

		dbRes := db.SaveExcerciese(dbExcerciese, pool)
		http_res := matcher.MatchResult(dbRes)
		writer.WriteHeader(http_res.GetStatus())
		writer.Write(http_res.GetData())
	})
}

// func UploadExcercieseHandler(pool *pgx.ConnPool) http.HandlerFunc {
// 	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

// 		if request.Method != "POST" {
// 			http.Error(writer, "Unsupported method", 405)
// 			return
// 		}
// 		// move to auth
// 		cookie, err := request.Cookie("bmstuOlympAuth")
// 		// unauth
// 		if err != nil {
// 			http.Error(writer, "Unauthorized", 403)
// 			return
// 		}

// 		auth_res := auth.AuthUser(cookie.Value)
// 		if auth_res.IsError() {
// 			return
// 		}
// 		// also here we got author id
// 		user_id, auth := auth.AuthUser(cookie.Value)
// 		// if !auth {
// 		// 	http.Error(writer, "Unauthorized", 403)
// 		// 	return
// 		// }

// 		if request.Body == nil {
// 			http.Error(writer, "Please send a request body", 400)
// 			return
// 		}
// 		body, err := ioutil.ReadAll(request.Body)
// 		defer request.Body.Close()
// 		if err != nil {
// 			http.Error(writer, "Please send a request body", 400)
// 			return
// 		}

// 		var excerciese views.ExcercieseView
// 		err = json.Unmarshal(body, &excerciese)

// 		if err != nil {
// 			http.Error(writer, "Error json", 400)
// 			return
// 		}

// 		excercieseEntity := excerciese.ToEntity()
// 		excercieseEntity.SetAuthor(user_id)

// 		file, err := base64.StdEncoding.DecodeString(excerciese.FileBase64)
// 		if err != nil {
// 			http.Error(writer, "Incorrect file", 400)
// 			return
// 		}

// 		// represent name in file storage
// 		newName := fstorage.ComputeName(excerciese.FileName)

// 		excercieseEntity.SetFileName(newName)
// 		err = db.SaveExcerciese(excercieseEntity, pool)
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}

// 		err = fstorage.WriteFile(file, newName, ".pdf")
// 		if err != nil {
// 			http.Error(writer, "Error save file", 500)
// 			return
// 		}

// 		err := sender.SendAnswer(0, "hi")
// 		if err := nil {
// 			// db.RemoveExcerciese(excercieseEntity.Id)
// 			// fs.RemoveFile(newName)
// 			return
//  		}

// 		writer.Write([]byte("SUCCESS")) //
// 	})
// }

// func GetExcerciese(pool *pgx.ConnPool) http.HandlerFunc {
// 	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

// 		if request.Method != "GET" {
// 			http.Error(writer, "Unsopported method", 405)
// 			return
// 		}

// 		// Get path variable from path
// 		idStr := strings.TrimPrefix(request.URL.Path, "/api/excercieses/get/")
// 		id, err := strconv.Atoi(idStr)
// 		if err != nil {
// 			http.Error(writer, "Incorrect path variable", http.StatusBadRequest)
// 		}
// 		uId := uint(id)

// 		res := db.GetExcerciese(uId, pool)
// 		httpRes := matcher.MatchResult(res)
// 		writer.WriteHeader(httpRes.GetStatus())
// 		writer.Write(httpRes.GetData())
// 	})
// }

// func GetExcercieses(pool *pgx.ConnPool) http.HandlerFunc {
// 	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
// 		if request.Method != "GET" {
// 			http.Error(writer, "Unsopported method", 405)
// 			return
// 		}

// 		query := request.URL.Query()
// 		pathVariablesStr := strings.TrimPrefix(request.URL.Path, "/api/excercieses/list/")
// 		vars := strings.Split(pathVariablesStr, "/")
// 		subject := ""
// 		tag := ""
// 		level := -1

// 		if len(vars) == 0 {
// 			http.Error(writer, "Not enough parameter", 404)
// 			return
// 		}

// 		// its very scary !!!!!
// 		for i, data := range vars {
// 			if i == 0 {
// 				subject = vars[i]
// 			}
// 			if i == 1 {
// 				tag = vars[i]
// 			}
// 			if i == 2 && data != "" {
// 				var err error
// 				level, err = strconv.Atoi(data)
// 				if err != nil {
// 					http.Error(writer, "INCORRECT PATH", 404)
// 				}
// 			}
// 		}

// 		limitArr := query["limit"]
// 		limit := -1
// 		if len(limitArr) > 0 {
// 			limit, _ = strconv.Atoi(limitArr[0])
// 		}
// 		offsetArr := query["offset"]
// 		offset := -1
// 		if len(offsetArr) > 0 {
// 			offset, _ = strconv.Atoi(offsetArr[0])
// 		}

// 		// its fucking crutch maybe, todo refactor !!!!!!!!!

// 		// check order for quer
// 		order := query["order"]
// 		is_desc := false
// 		if len(order) > 0 && order[0] == "desc" {
// 			is_desc = true
// 		}

// 		// 1 - subject 2 - tag 3 - level
// 		// query 1 - limit 2 - offset 3 - order

// 		res := db.GetExcercieseList(tag, subject, level, limit, offset, is_desc, pool)
// 		httpRes := matcher.MatchResult(res)
// 		writer.WriteHeader(httpRes.GetStatus())
// 		writer.Write([]byte(val))
// 	})
// }
