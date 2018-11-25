package controllers

import (
	"net/http"
	"strings"

	"github.com/OlympBMSTU/exercises/db"
)

// GetSubjectsHandler : returns all olympyad subjects
func GetSubjectsHandler(writer http.ResponseWriter, request *http.Request) {
	userID := CheckMethodAndAuthenticate(writer, request, "GET")
	if userID == nil {
		return
	}

	dbRes := db.GetSubjects(request.Context())
	WriteResponse(&writer, "JSON", dbRes)
}

// GetTagsHandler : its return all tags, thats references to this subject
func GetTagsHandler(writer http.ResponseWriter, request *http.Request) {
	userID := CheckMethodAndAuthenticate(writer, request, "GET")
	if userID == nil {
		return
	}

	subject := strings.TrimPrefix(request.URL.Path, "/api/exercises/tags/")
	dbRes := db.GetTgasBySubect(subject, request.Context())

	WriteResponse(&writer, "JSON", dbRes)
}
