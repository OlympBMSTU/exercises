package views

import "github.com/OlympBMSTU/exercises/entities"

type ExerciseView struct {
	ID       int      `json:"id"`
	Answer   string   `json:"answer"`
	FileName string   `json:"file_name"`
	Subject  string   `json:"subject"`
	Tags     []string `json:"tags"`
	Level    int      `json:"level"`
	IsBroken bool     `json:"is_broken"` /// ???????
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

// TODO how correctly check
func (view ExerciseView) IsEmpty() bool {
	return view.Answer == "" && view.Subject == "" && len(view.Tags) == 0 && view.Level == -1 && view.IsBroken == false
}

func (view *ExerciseView) SetFileName(fileName string) {
	view.FileName = fileName
}

func (view *ExerciseView) SetAuthor(author uint) {
	view.Author = author
}
