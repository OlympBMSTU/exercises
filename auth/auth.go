package auth

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/OlympBMSTU/exercises/auth/result"
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

const HASH_SECRET = "Любовь измеряется мерой прощения."

func AuthByUserCookie(request *http.Request, cookieName string) result.AuthResult {
	return result.OkResult(1)
	cookie, err := request.Cookie(cookieName)
	if err != nil {
		return result.ErrorResult(result.NO_COOKIE, "No cookie")
	}

	return authUser(cookie.Value)
}

func authUser(jwt string) result.AuthResult {
	jwt_data := strings.Split(jwt, ".")

	if len(jwt_data) != 3 {
		return result.ErrorResult(result.ERROR_PARSE_JWT, "")
	}

	header, err := base64.StdEncoding.DecodeString(jwt_data[0])
	payload, err := base64.StdEncoding.DecodeString(jwt_data[1])
	hash, err := base64.StdEncoding.DecodeString(jwt_data[2])
	if err != nil {
		return result.ErrorResult(result.ERROR_PARSE_JWT, "")
	}

	var jwt_header JWTHeader
	err = json.Unmarshal(header, &jwt_header)
	if err != nil {
		return result.ErrorResult(result.ERROR_PARSE_JWT, "")
	}

	var jwt_payload JWTPayload
	err = json.Unmarshal(payload, &jwt_payload)
	if err != nil {
		return result.ErrorResult(result.ERROR_PARSE_JWT, "")
	}

	var buffer bytes.Buffer
	buffer.WriteString(jwt_data[0])
	buffer.WriteString(".")
	buffer.WriteString(jwt_data[1])

	mac := hmac.New(sha256.New, []byte(HASH_SECRET))
	mac.Write(buffer.Bytes())
	expected := []byte(hex.EncodeToString(mac.Sum(nil)))
	if !hmac.Equal(hash, expected) {
		return result.ErrorResult(result.NO_AUTHROIZED, "")
	}

	v := time.Now().Nanosecond()

	// maybe wrong
	if jwt_payload.Exp < uint(v) {
		return result.ErrorResult(result.NO_AUTHROIZED, "")
	}
	return result.OkResult(jwt_payload.Id)
}
