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
	return SenderResult{}
}

func ErrorResult(params ...interface{}) SenderResult {
	return SenderResult{}
}
