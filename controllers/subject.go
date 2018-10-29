package controllers

import (
	"net/http"
	"strings"

	"github.com/OlympBMSTU/exercises/auth"
	"github.com/OlympBMSTU/exercises/db"
	"github.com/jackc/pgx"
)

func GetSubjects(pool *pgx.ConnPool) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		OptionsCredentials(&writer)
		if request.Method == "OPTIONS" {
			writer.Write([]byte("hi"))
			return
		}

		if request.Method != "GET" {
			http.Error(writer, "Unsupported method", 405)
			return
		}
		writer.Header().Set("Content-Type", "application/json")

		authRes := auth.AuthByUserCookie(request, "bmstuOlympAuth")
		if authRes.IsError() {
			WriteResponse(&writer, authRes)
			return
		}

		dbRes := db.GetSubjects(pool)
		WriteResponse(&writer, dbRes)
	})
}

func GetTags(pool *pgx.ConnPool) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		OptionsCredentials(&writer)
		if request.Method == "OPTIONS" {
			writer.Write([]byte("hi"))
			return
		}
		if request.Method != "GET" {
			http.Error(writer, "Unsupported method", 405)
			return
		}
		writer.Header().Set("Content-Type", "application/json")

		authRes := auth.AuthByUserCookie(request, "bmstuOlympAuth")
		if authRes.IsError() {
			WriteResponse(&writer, authRes)
			return
		}

		subject := strings.TrimPrefix(request.URL.Path, "/api/exercises/tags/")
		dbRes := db.GetTgasBySubect(subject, pool)

		WriteResponse(&writer, dbRes)
	})
}
