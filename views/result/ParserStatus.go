package result

type ParserStatus struct {
	code  int
	descr string
}

func (status ParserStatus) GetCode() int {
	return status.code
}

func (status ParserStatus) GetDescription() string {
	return status.descr
}

func (status ParserStatus) IsError() bool {
	return false
}
