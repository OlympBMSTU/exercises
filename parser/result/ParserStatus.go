package result

const (
	NO_ERROR                  = 0
	INCORRECT_BODY            = 1
	INCORRECT_TAGS            = 2
	INCORRECT_LEVEL           = 3
	INCORRECT_ANSWER_ARR      = 4
	INCORRECT_MARK            = 5
	INCORRECT_TYPE_OLYMP      = 6
	INCORRECT_CLASS           = 7
	INCORRECT_POSITION        = 8
	INCORRECT_MARK_IN_ANSWERS = 9
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
	return !(status.code == NO_ERROR)
}

// TODO
func parseError(err error) {
}
