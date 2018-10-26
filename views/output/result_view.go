package output

type ResultView struct {
	Data    interface{}
	Message string
	Status  string
}

func (res ResultView) GetData() interface{} {
	return res.Data
}

func (res *ResultView) SetData(data interface{}) {
	res.Data = data
}

func (res *ResultView) GetMessage() string {
	return res.Message
}

func (res *ResultView) SetMessage(message string) {
	res.Message = message
}

func (res *ResultView) GetStatus() string {
	return res.Status
}

func (res *ResultView) SetStatus(status string) {
	res.Status = status
}

// json: {
// "status": Ok,
// "message": error,
// "data": data
// }
