package result

const (
	NO_ERROR        = 0
	ERROR_PARSE_JWT = 1
	NO_AUTHROIZED   = 2
)

type AuthStatus struct {
	code  int
	descr string
}

func (status AuthStatus) IsError() bool {
	return !(status.code == NO_ERROR)
}

func (status AuthStatus) GetCode() int {
	return status.code
}

func (status AuthStatus) GetDescription() string {
	return status.descr
}

func parseError(err error) AuthStatus {
	return AuthStatus{}
}
