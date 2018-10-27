package controllers

import (
	"net/http"

	matcher "github.com/OlympBMSTU/excericieses/controllers/matcher_result"
	"github.com/OlympBMSTU/excericieses/result"
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
