package controllers

import (
	"net/http"
	"strings"

	matcher "github.com/OlympBMSTU/excericieses/controllers/matcher_result"
	"github.com/OlympBMSTU/excericieses/db"
	"github.com/OlympBMSTU/excericieses/result"
	"github.com/jackc/pgx"
)

func GetSubjects(pool *pgx.ConnPool) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != "GET" {
			http.Error(writer, "Unsupported method", 405)
			return
		}
		writer.Header().Set("Content-Type", "application/json")

		var res result.Result
		res = db.GetSubjects(pool)
		httpRes := matcher.MatchResult(res)

		writer.WriteHeader(httpRes.GetStatus())
		writer.Write(httpRes.GetData())
	})
}

func GetTags(pool *pgx.ConnPool) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != "GET" {
			http.Error(writer, "Unsupported method", 405)
			return
		}
		writer.Header().Set("Content-Type", "application/json")

		subject := strings.TrimPrefix(request.URL.Path, "/api/excercises/tags/")
		res := db.GetTgasBySubect(subject, pool)

		httpRes := matcher.MatchResult(res)
		writer.WriteHeader(httpRes.GetStatus())
		writer.Write(httpRes.GetData())
	})
}
