package result

// TODO return data not interface but structure
// add method rew  data maybe
type Result interface {
	IsError() bool
	GetStatus() Status
	GetData() Data
}
