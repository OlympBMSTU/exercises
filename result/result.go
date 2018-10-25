package result

type Result interface {
	IsError() bool
	GetStatus() Status
	GetData() interface{}
}
