package result

import "github.com/OlympBMSTU/exercises/result"

type ParserResult struct {
	data   ParserData
	status ParserStatus
}

func (res ParserResult) GetData() interface{} {
	return res.data.GetData()
}

func (res ParserResult) IsError() bool {
	return res.status.IsError()
}

func (res ParserResult) GetStatus() result.Status {
	return res.status
}

func ErrorResult(params ...interface{}) ParserResult {
	return ParserResult{}
}

func OkResult(data interface{}) ParserResult {
	return ParserResult{}
}
