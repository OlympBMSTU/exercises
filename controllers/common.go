package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/OlympBMSTU/exercises/auth"
	matcher "github.com/OlympBMSTU/exercises/controllers/matcher_result"
	"github.com/OlympBMSTU/exercises/result"
)

func writeResponseFromResult(writer *http.ResponseWriter, format string, res result.Result) {
	httpRes := matcher.MatchResult(res)
	(*writer).WriteHeader(httpRes.GetStatus())
	(*writer).Write(httpRes.GetData())
}

// TODO for xml, other formats
func writeResponseFromMap(writer *http.ResponseWriter, format string, data map[string]interface{}, status int) {
	val, err := json.Marshal(data)
	if err != nil {
		(*writer).WriteHeader(http.StatusInternalServerError)
		(*writer).Write([]byte("Internal server error: cant create json"))
		return
	}

	(*writer).WriteHeader(status)
	(*writer).Write(val)
}

func WriteResponse(writer *http.ResponseWriter, format string, params ...interface{}) {
	if len(params) < 2 {
		writeResponseFromResult(writer, format, params[0].(result.Result))
	} else {
		writeResponseFromMap(writer, format, params[0].(map[string]interface{}), params[1].(int))
	}
}

func OptionsCredentials(writer *http.ResponseWriter) {
	(*writer).Header().Set("Access-Control-Allow-Method", "POST, OPTIONS")
	(*writer).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	(*writer).Header().Set("Access-Control-Allow-Headers", "Content-Type, access-controll-request-method,x-requested-with")
	(*writer).Header().Set("Access-Control-Allow-Credentials", "true")
}

func checkMethod(writer http.ResponseWriter, req *http.Request, method string) bool {
	OptionsCredentials(&writer)
	if req.Method == "OPTIONS" {
		writer.Write([]byte("hi"))
		return false
	}
	writer.Header().Set("Content-Type", "application/json")

	if req.Method != method {
		WriteResponse(&writer, "JSON", map[string]interface{}{
			"Message": "Unsupported method",
			"Status":  "Error",
			"Data":    nil,
		}, http.StatusMethodNotAllowed)
		return false
	}

	return true
}

func authenticateUser(writer http.ResponseWriter, req *http.Request) *uint {
	authRes := auth.AuthByUserCookie(req)
	if authRes.IsError() {
		WriteResponse(&writer, "JSON", authRes)
		return nil
	}
	userID := authRes.GetData().(uint)
	return &userID
}

func CheckMethodAndAuthenticate(writer http.ResponseWriter, req *http.Request, method string) *uint {
	if !checkMethod(writer, req, method) {
		return nil
	}

	return authenticateUser(writer, req)
}
