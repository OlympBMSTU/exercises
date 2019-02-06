package auth

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/OlympBMSTU/exercises/auth/result"
	"github.com/OlympBMSTU/exercises/config"
	"github.com/OlympBMSTU/exercises/logger"
)

type JWTHeader struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type JWTPayload struct {
	Id   uint   `json:"id"`
	Type uint   `json:"type"`
	Iss  string `json:"iss"`
	Sub  string `json:"sub"`
	Exp  uint   `json:"exp"`
}

func AuthByUserCookie(request *http.Request) result.AuthResult {
	conf, _ := config.GetConfigInstance()
	if conf.IsTest() {
		return result.OkResult(1)
	}

	cookieName := conf.GetAuthCookieName()
	log := logger.GetLogger()
	log.Info(fmt.Sprintf("cookieName: %s", cookieName))

	cookie, err := request.Cookie(cookieName)
	if err != nil {
		log.Error("AuthError", errors.New("Cookie missing"))
		return result.ErrorResult(result.NO_COOKIE, "No cookie")
	}

	return authUser(cookie.Value, conf.GetAuthSecret())
}

func authUser(jwt string, hashSecret string) result.AuthResult {
	log := logger.GetLogger()
	jwt_norm, err := url.QueryUnescape(jwt)
	if err != nil {
		log.Error("AuthError", err)
		return result.ErrorResult(result.ERROR_PARSE_JWT, "")
	}

	jwt_data := strings.Split(jwt_norm, ".")

	if len(jwt_data) != 3 {
		log.Error("AuthError", errors.New("JWT len is not equal 3"))
		return result.ErrorResult(result.ERROR_PARSE_JWT, "")
	}

	header, err := base64.StdEncoding.DecodeString(jwt_data[0])
	if err != nil {
		log.Error("AuthError", err)
		return result.ErrorResult(result.ERROR_PARSE_JWT, "")
	}

	payload, err := base64.StdEncoding.DecodeString(jwt_data[1])
	if err != nil {
		log.Error("AuthError", err)
		return result.ErrorResult(result.ERROR_PARSE_JWT, "")
	}

	hash, err := base64.StdEncoding.DecodeString(jwt_data[2])
	if err != nil {
		log.Error("AuthError", err)
		return result.ErrorResult(result.ERROR_PARSE_JWT, "")
	}

	var jwt_header JWTHeader
	err = json.Unmarshal(header, &jwt_header)
	if err != nil {
		log.Error("AuthError", err)
		return result.ErrorResult(result.ERROR_PARSE_JWT, "")
	}

	var jwt_payload JWTPayload
	err = json.Unmarshal(payload, &jwt_payload)
	if err != nil {
		log.Error("AuthError", err)
		return result.ErrorResult(result.ERROR_PARSE_JWT, "")
	}

	var buffer bytes.Buffer
	buffer.WriteString(jwt_data[0])
	buffer.WriteString(".")
	buffer.WriteString(jwt_data[1])

	mac := hmac.New(sha256.New, []byte(hashSecret))
	mac.Write(buffer.Bytes())
	expected := []byte(hex.EncodeToString(mac.Sum(nil)))
	if !hmac.Equal(hash, expected) {
		log.Error("AuthError", errors.New("Not equal cookie"))
		return result.ErrorResult(result.NO_AUTHROIZED, "")
	}

	v := time.Now().Nanosecond()

	// maybe wrong
	if jwt_payload.Exp < uint(v) {
		log.Error("AuthError", errors.New("Cookie is expired"))
		return result.ErrorResult(result.NO_AUTHROIZED, "")
	}
	return result.OkResult(jwt_payload.Id)
}
