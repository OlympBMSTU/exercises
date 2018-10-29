package result

const (
	NO_ERROR          = 0
	ERROR_SEND        = 1
	ERROR_CREATE_JSON = 2
)

type SenderStatus struct {
	code  int
	descr string
}

func (status SenderStatus) GetCode() int {
	return status.code
}

func (status SenderStatus) GetDescription() string {
	return status.descr
}

func (status SenderStatus) IsError() bool {
	return !(status.code == NO_ERROR)
}
