package controllers

import (
	"net/http"
	"strings"

	"github.com/OlympBMSTU/exercises/db"
)

func GetSubjects(writer http.ResponseWriter, request *http.Request) {
	userId := CheckMethodAndAuthenticate(writer, request, "GET")
	if userId == nil {
		return
	}

	dbRes := db.GetSubjects(request.Context())
	WriteResponse(&writer, "JSON", dbRes)
}

func GetTags(writer http.ResponseWriter, request *http.Request) {
	userId := CheckMethodAndAuthenticate(writer, request, "GET")
	if userId == nil {
		return
	}

	subject := strings.TrimPrefix(request.URL.Path, "/api/exercises/tags/")
	dbRes := db.GetTgasBySubect(subject, request.Context())

	WriteResponse(&writer, "JSON", dbRes)
}
