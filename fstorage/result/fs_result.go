package result

import (
	"github.com/OlympBMSTU/excericieses/result"
)

type FSResult struct {
	data   FSData
	status FSStatus
}

func (res FSResult) GetData() interface{} {
	return res.data.GetData()
}

func (res FSResult) GetStatus() result.Status {
	return res.status
}

func (res FSResult) IsError() bool {
	return res.status.IsError()
}

func OkResult(data interface{}) FSResult {
	return FSResult{
		FSData{data},
		FSStatus{NO_ERROR, ""},
	}
}

func ErrorResult(err error) FSResult {
	return FSResult{
		FSData{nil},
		FSStatus{ERROR_CREATE_FILE, ""},
	}
}
