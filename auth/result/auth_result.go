package result

type AuthResult struct {
	data   AuthData
	status AuthStatus
}

func ErrroResult(params ...interface{}) AuthResult {
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

func OkResult(id uint) AuthResult {
	return AuthResult{
		AuthData{id},
		AuthStatus{NO_ERROR, ""},
	}
}
