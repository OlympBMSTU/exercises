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
	return entities.ExerciseEntity{
		AuthorId: view.Author,
		FileName: view.FileName,
		Tags:     view.Tags,
		Level:    view.Level,
		Subject:  view.Subject,
		IsBroken: false,
	}
}

func (view *ExerciseView) SetFileName(fileName string) {
	view.FileName = fileName
}

func (view *ExerciseView) SetAuthor(author uint) {
	view.Author = author
}
