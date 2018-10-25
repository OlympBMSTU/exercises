package auth

const (
	NO_ERROR        = 0
	ERROR_PARSE_JWT = 1
	NO_AUTHROIZED   = 2
)

type AuthResult struct {
	data   AuthData
	status AuthStatus
}

type AuthData struct {
	data interface{}
}

type AuthStatus struct {
	code  int
	descr string
}

func parseError(err error) AuthStatus {
	return AuthStatus{}
}

func errroResult(params ...interface{}) AuthResult {
	if len(params) == 1 {
		return AuthResult{
			AuthData{nil},
			parseError(params[0].(error)),
		}
	} else {
		return AuthResult{
			AuthData{nil},
			AuthStatus{
				params[0].(int),
				params[1].(string),
			},
		}
	}
}

func okResult(id uint) AuthResult {
	return AuthResult{
		AuthData{id},
		AuthStatus{NO_ERROR, ""},
	}
}
