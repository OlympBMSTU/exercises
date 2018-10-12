package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/OlympBMSTU/excericieses/db"
	"github.com/jackc/pgx"
)

func GetSubjects(pool *pgx.ConnPool) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != "GET" {
			http.Error(writer, "Unsupported method", 404)
			return
		}

		subs, err := db.GetSubjects(pool)
		if err != nil {
			http.Error(writer, "Internal server error", http.StatusInternalServerError)
			return
		}

		val, err := json.Marshal(subs)
		if err != nil {
			http.Error(writer, "Internal server error", http.StatusInternalServerError)
			return
		}

		writer.Write([]byte(val))
	})
}

func GetTags(pool *pgx.ConnPool) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != "GET" {
			http.Error(writer, "Unsupported method", 404)
			return
		}

		// var subject string
		// if len(request.URL.Query()["subject"]) < 1 {
		// 	http.Error(writer, "Incorrect request", 403)
		// 	return
		// }

		// subject = request.URL.Query()["subject"][0]

		subject := strings.TrimPrefix(request.URL.Path, "/api/excercieses/tags/")
		// str = strings.TrimRight(str, "/tags")
		fmt.Print(subject)
		tags, err := db.GetTgasBySubect(subject, pool)
		if err != nil {
			http.Error(writer, "Internal server error", 500)
			return
		}

		val, err := json.Marshal(tags)
		if err != nil {
			http.Error(writer, "Internal server error", 500)
			return
		}

		writer.Write([]byte(val))
	})
}
