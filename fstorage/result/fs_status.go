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
