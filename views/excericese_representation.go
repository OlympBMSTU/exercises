package views

type ExcercieseRepresentation struct {
	Id       uint     `json:"id"`
	Answer   string   `json:"answer"`
	FileName string   `json:"file_name"`
	Subject  string   `json:"subject"`
	Tags     []string `json:"tags"`
	Level    uint     `json:"level"`
	Author   uint
}
