package controllers

import (
	"net/http"
	"strings"

	"github.com/OlympBMSTU/exercises/db"
)

func GetSubjects(writer http.ResponseWriter, request *http.Request) {
	if !CheckMethodAndAuthenticate(writer, request, "GET") {
		return
	}

	dbRes := db.GetSubjects(request.Context())
	WriteResponse(&writer, "JSON", dbRes)
}

func GetTags(writer http.ResponseWriter, request *http.Request) {
	if !CheckMethodAndAuthenticate(writer, request, "GET") {
		return
	}

	subject := strings.TrimPrefix(request.URL.Path, "/api/exercises/tags/")
	dbRes := db.GetTgasBySubect(subject, request.Context())

	WriteResponse(&writer, "JSON", dbRes)
}
