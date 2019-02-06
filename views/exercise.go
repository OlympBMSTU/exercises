package views

import "github.com/OlympBMSTU/exercises/entities"

// class integer,
// position INTEGER,
// mark   INTEGER,
// type_olymp INTEGER,
// answer jsonb,
// created timestamp default NOW(

type ExerciseView struct {
	ID        int          `json:"id"`
	Answer    string       `json:"answer"`
	FileName  string       `json:"file_name"`
	Subject   string       `json:"subject"`
	Tags      []string     `json:"tags"`
	Level     int          `json:"level"`
	IsBroken  bool         `json:"is_broken"` /// ???????
	Class     int          `json:"class"`
	Position  int          `json"position"`
	Mark      int          `json"mark"`
	TypeOlymp int          `json"type_olymp"`
	Answers   []AnswerView `json:"answers"`
	Author    uint
}

func (view ExerciseView) ToExEntity() entities.ExerciseEntity {
	answEntities := make([]entities.Answer, len(view.Answer))
	for _, ans := range view.Answers {
		// no need
		answEntities = append(answEntities, entities.Answer{
			ID:     ans.ID,
			Input:  ans.Input,
			Output: ans.Output,
		})
	}
	return entities.ExerciseEntity{
		Id:        uint(view.ID),
		AuthorId:  view.Author,
		FileName:  view.FileName,
		Tags:      view.Tags,
		Level:     view.Level,
		Subject:   view.Subject,
		IsBroken:  view.IsBroken,
		Class:     view.Class,
		Mark:      view.Mark,
		Position:  view.Position,
		TypeOlymp: view.TypeOlymp,
		Answers:   answEntities,
	}
}

// TODO how correctly check
func (view ExerciseView) IsEmpty() bool {
	return (len(view.Answers) == 0) && view.Subject == "" && len(view.Tags) == 0 && view.Level == -1 && view.IsBroken == false
}

func (view *ExerciseView) SetFileName(fileName string) {
	view.FileName = fileName
}

func (view *ExerciseView) SetAuthor(author uint) {
	view.Author = author
}
