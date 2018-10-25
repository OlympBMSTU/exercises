package result

type FSData struct {
	data interface{}
}

func (dt FSData) GetData() interface{} {
	return dt.data
}
