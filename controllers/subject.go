package controllers

import (
	"net/http"
	"strings"

	"github.com/OlympBMSTU/exercises/auth"
	"github.com/OlympBMSTU/exercises/db"
)

func GetSubjects(writer http.ResponseWriter, request *http.Request) {
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

	dbRes := db.GetSubjects(request.Context())
	WriteResponse(&writer, dbRes)
}

func GetTags(writer http.ResponseWriter, request *http.Request) {
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
	dbRes := db.GetTgasBySubect(subject, request.Context())

	WriteResponse(&writer, dbRes)
}
