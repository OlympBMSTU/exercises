package views

import "github.com/OlympBMSTU/exercises/entities"

type ExerciseView struct {
	Answer   string   `json:"answer"`
	FileName string   `json:"file_name"`
	Subject  string   `json:"subject"`
	Tags     []string `json:"tags"`
	Level    uint     `json:"level"`
	Author   uint
}

func (view ExerciseView) ToExEntity() entities.ExerciseEntity {
	return entities.ExerciseEntity{}
}
