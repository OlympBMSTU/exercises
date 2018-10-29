package result

import (
	"github.com/OlympBMSTU/exercises/result"
)

type SenderResult struct {
	data   SenderData
	status SenderStatus
}

func (res SenderResult) GetData() interface{} {
	return res.data.GetData()
}

func (res SenderResult) GetStatus() result.Status {
	return res.status
}

func (res SenderResult) IsError() bool {
	return res.status.IsError()
}

func OkResult(params ...interface{}) SenderResult {
	return SenderResult{
		SenderData{nil},
		SenderStatus{NO_ERROR, ""},
	}
}

func ErrorResult(params ...interface{}) SenderResult {
	if len(params) < 2 {
		return SenderResult{
			SenderData{nil},
			SenderStatus{ERROR_SEND, ""},
		}
	}
	return SenderResult{
		SenderData{nil},
		SenderStatus{params[0].(int), params[1].(string)},
	}
}
