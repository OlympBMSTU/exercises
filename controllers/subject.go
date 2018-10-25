package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/OlympBMSTU/excericieses/db"
	"github.com/OlympBMSTU/excericieses/result"
	"github.com/jackc/pgx"
)

func GetSubjects(pool *pgx.ConnPool) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != "GET" {
			http.Error(writer, "Unsupported method", 404)
			return
		}

		var res result.Result
		res = db.GetSubjects(pool)
		val, err := json.Marshal(res.GetData())

		// res := db.GetSubjects(pool)
		// res.GetResponseData() // структура подобна http response
		// res.Data, res.Code, res.Descr
		// маршиализация там же
		// Если маршализация отвалилась то возвр 500 Internal server error

		// res.unwrap()
		// if res.Error() {
		// 	err := res.GetError()

		// 	descr, code = res.GetHttpResponse() //ErrorMatcher(err)

		// }

		//  := db.GetSubjects(pool)
		// if err != nil {
		// 	http.Error(writer, "Internal server error", http.StatusInternalServerError)
		// 	return
		// }

		// val, err := json.Marshal(subs)
		if err != nil {
			http.Error(writer, "Internal server error", http.StatusInternalServerError)
			return
		}

		// writer.WriteHeader

		writer.Write([]byte(val))
	})
}

// func GetTags(pool *pgx.ConnPool) http.HandlerFunc {
// 	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
// 		if request.Method != "GET" {
// 			http.Error(writer, "Unsupported method", 404)
// 			return
// 		}

// 		subject := strings.TrimPrefix(request.URL.Path, "/api/excercieses/tags/")
// 		// str = strings.TrimRight(str, "/tags")
// 		fmt.Print(subject)
// 		tags, err := db.GetTgasBySubect(subject, pool)
// 		if err != nil {
// 			http.Error(writer, "Internal server error", 500)
// 			return
// 		}

// 		val, err := json.Marshal(tags)
// 		if err != nil {
// 			http.Error(writer, "Internal server error", 500)
// 			return
// 		}

// 		writer.Write([]byte(val))
// 	})
// }
