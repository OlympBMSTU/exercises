package fstorage

const (
	NO_ERROR          = 0
	ERROR_CREATE_FILE = 1
	ERROR_CREATE_DIRS = 2
	ERROR_WRITE_FILE  = 3
)

type FStorageError struct {
	code  int
	descr string
}

type FStagereRes struct {
	err  FStorageError
	data interface{}
}
