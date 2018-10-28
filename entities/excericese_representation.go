package entities

type ExerciseRepresentation struct {
	Id       uint     `json:"id"`
	FileName string   `json:"file_name"`
	Subject  string   `json:"subject"`
	Tags     []string `json:"tags"`
	Level    uint     `json:"level"`
	Author   uint     `json:"author"`
}

func NewExRepresentation(ex ExerciseEntity) ExerciseRepresentation {
	return ExerciseRepresentation{
		Id:       ex.Id,
		FileName: ex.FileName,
		Subject:  ex.Subject,
		Tags:     ex.Tags,
		Level:    ex.Level,
		Author:   ex.AuthorId,
	}
}