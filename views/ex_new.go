package views

type ExcercieseViewN struct {
	answer   string   `json:"answer"`
	fileName string   `json:"file_name"`
	subject  string   `json:"subject"`
	tags     []string `json:"tags"`
	level    uint     `json:"level"`
	author   uint
}

func (view ExcercieseViewN) IsValid() bool {
	return true
}
