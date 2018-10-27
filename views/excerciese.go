package views

import "github.com/OlympBMSTU/excericieses/entities"

type ExcercieseView struct {
	Answer     string   `json:"answer"`
	FileBase64 string   `json:"file"`
	FileName   string   `json:"file_name"`
	Subject    string   `json:"subject"`
	Tags       []string `json:"tags"`
	Level      uint     `json:"level"`
	Author     uint
}

// func (view *ExcercieseView) ToEntity() entities.ExcercieseEntity {
// 	return entities.NewExcercieseEntity(
// 		view.Answer,
// 		view.Tags,
// 		view.Level,
// 		view.Subject,
// 	)
// }

func ExcercieseViewFrom(entity entities.ExcercieseEntity, tags []string) ExcercieseView {
	return ExcercieseView{
		entity.GetRightAnswer(),
		"",
		entity.GetFileName(),
		entity.GetSubject(),
		tags,
		entity.GetLevel(),
		entity.GetAuthorId(),
	}
}
