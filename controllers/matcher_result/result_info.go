package matcher_result

const (
	statusOK           = "OK"
	statusError        = "Error"
	statusSubjectError = "SebjectErr"
)

type ResultInfo struct {
	Message  string
	HttpCode int
	Status   string
}

func NewResultInfo(message string, code int, status string) ResultInfo {
	return ResultInfo{
		Message:  message,
		HttpCode: code,
		Status:   status,
	}
}
