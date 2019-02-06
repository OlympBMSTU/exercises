package result

type AuthData struct {
	data interface{}
}

func (data AuthData) GetData() interface{} {
	return data.data
}
