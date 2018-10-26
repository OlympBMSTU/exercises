package controllers

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
