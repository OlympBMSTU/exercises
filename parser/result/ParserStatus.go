package result

const (
	NO_ERROR        = 0
	INCORRECT_BODY  = 1
	INCORRECT_TAGS  = 2
	INCORRECT_LEVEL = 3
)

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
