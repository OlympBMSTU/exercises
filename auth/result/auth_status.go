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

func parseError(err error) AuthStatus {
	return AuthStatus{}
}
