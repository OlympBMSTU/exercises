package result

// todo data as View base
type ParserData struct {
	data interface{}
}

func (dat ParserData) GetData() interface{} {
	return dat.data
}
