package result

import "github.com/OlympBMSTU/exercises/result"

type AuthResult struct {
	data   AuthData
	status AuthStatus
}

func (res AuthResult) GetData() interface{} {
	return res.data.GetData()
}

func (res AuthResult) IsError() bool {
	return res.status.IsError()
}

func (res AuthResult) GetStatus() result.Status {
	return res.status
}

func ErrorResult(params ...interface{}) AuthResult {
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
