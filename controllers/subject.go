package controllers

import (
	"encoding/json"
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

type SubjecJson struct {
	Name string `json:"name"`
}

func CreateSubectHandler(writer http.ResponseWriter, request *http.Request) {
	userID := CheckMethodAndAuthenticate(writer, request, "POST")
	if userID == nil {
		return
	}

	decoder := json.NewDecoder(request.Body)
	var subj SubjecJson
	err := decoder.Decode(&subj)
	if err != nil {
		WriteResponse(&writer, "JSON",
			map[string]interface{}{
				"Message": "Error parse json subject",
				"Data":    nil,
				"Status":  "Error",
			}, http.StatusBadRequest)
		return
	}

	dbRes := db.SaveSubject(subj.Name, request.Context())
	WriteResponse(&writer, "JSON", dbRes)
}
