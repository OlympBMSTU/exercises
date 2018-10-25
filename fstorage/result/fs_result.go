package result

type FSResult struct {
	data   FSData
	status FSStatus
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
