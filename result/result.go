package result

type Result interface {
	IsError() bool
	GetData() interface{}
}
