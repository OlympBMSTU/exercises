package result

import "github.com/OlympBMSTU/exercises/result"

type ParserResult struct {
	data   ParserData
	status ParserStatus
}

func (res ParserResult) GetData() result.Data {
	return res.data
}

func (res ParserResult) IsError() bool {
	return res.status.IsError()
}

func (res ParserResult) GetStatus() result.Status {
	return res.status
}

// implement TODO call parse error
func ErrorResult(params ...interface{}) ParserResult {
	if len(params) < 2 {
		return ParserResult{
			ParserData{nil},
			ParserStatus{INCORRECT_BODY, "Incorrect body"},
		}
	}
	return ParserResult{
		ParserData{nil},
		ParserStatus{params[0].(int), params[1].(string)},
	}
}

func OkResult(data interface{}) ParserResult {
	return ParserResult{
		ParserData{data},
		ParserStatus{NO_ERROR, ""},
	}
}
