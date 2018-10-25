package result

type Status interface {
	IsError() bool
	GetCode() int
	GetDescription() string
}
