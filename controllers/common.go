package controllers

import (
	"net/http"

	"github.com/OlympBMSTU/exercises/auth"
	matcher "github.com/OlympBMSTU/exercises/controllers/matcher_result"
	"github.com/OlympBMSTU/exercises/result"
)

func WriteResponse(writer *http.ResponseWriter, res result.Result) {
	httpRes := matcher.MatchResult(res)
	(*writer).WriteHeader(httpRes.GetStatus())
	(*writer).Write(httpRes.GetData())
}

func OptionsCredentials(writer *http.ResponseWriter) {
	(*writer).Header().Set("Access-Control-Allow-Method", "POST, OPTIONS")
	(*writer).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	(*writer).Header().Set("Access-Control-Allow-Headers", "Content-Type, access-controll-request-method,x-requested-with")
	(*writer).Header().Set("Access-Control-Allow-Credentials", "true")
}

func CheckMethodAndAuthenticate(writer http.ResponseWriter, req *http.Request, method string) bool {
	OptionsCredentials(&writer)
	if req.Method == "OPTIONS" {
		writer.Write([]byte("hi"))
		return false
	}

	if req.Method != method {
		http.Error(writer, "Unsupported method", 405)
		return false
	}
	writer.Header().Set("Content-Type", "application/json")

	authRes := auth.AuthByUserCookie(req)
	if authRes.IsError() {
		WriteResponse(&writer, authRes)
		return false
	}
	return true
}
