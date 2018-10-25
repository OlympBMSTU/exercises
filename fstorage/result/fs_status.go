package result

const (
	NO_ERROR          = 0
	ERROR_CREATE_FILE = 1
	ERROR_CREATE_DIRS = 2
	ERROR_WRITE_FILE  = 3
)

type FSStatus struct {
	code  int
	descr string
}

func (status FSStatus) IsError() bool {
	return !(status.code == NO_ERROR)
}

func (status FSStatus) GetCode() int {
	return status.code
}

func (status FSStatus) GetDescription() string {
	return status.descr
}
