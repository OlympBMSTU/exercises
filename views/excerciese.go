package views

import "github.com/HustonMmmavr/excercieses/entities"

type ExcercieseView struct {
	Answer     string `json:"answer"`
	FileBase64 string `json:"file"`
	FileName   string `json:"file_name"`
	author     uint64
}

func (view *ExcercieseView) ToEntity() entities.ExcercieseEntity {
	return entities.ExcercieseEntity{}
}
