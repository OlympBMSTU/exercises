package entities

type ExerciseRepresentation struct {
	Id       uint     `json:"id"`
	FileName string   `json:"file_name"`
	Subject  string   `json:"subject"`
	Answer   string   `json:"answer"`
	Tags     []string `json:"tags"`
	Level    int      `json:"level"`
	Author   uint     `json:"author"`
	IsBroken bool     `json:"is_broken"`
}

func NewExRepresentation(ex ExerciseEntity) ExerciseRepresentation {
	return ExerciseRepresentation{
		Id:       ex.Id,
		FileName: ex.FileName,
		Answer:   "",
		Subject:  ex.Subject,
		Tags:     ex.Tags,
		Level:    ex.Level,
		Author:   ex.AuthorId,
		IsBroken: ex.IsBroken,
	}
}
