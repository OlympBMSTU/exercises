package entities

type ExerciseRepresentation struct {
	Id        uint     `json:"id"`
	FileName  string   `json:"file_name"`
	Subject   string   `json:"subject"`
	Tags      []string `json:"tags"`
	Level     int      `json:"level"`
	Author    uint     `json:"author"`
	IsBroken  bool     `json:"is_broken"`
	Class     int      `json:"class"`
	Position  int      `json:"position"`
	Mark      int      `json:"mark"`
	TypeOlymp int      `json:"type_olymp"`
	Created   string   `json:"created"`
}

func NewExRepresentation(ex ExerciseEntity) ExerciseRepresentation {
	return ExerciseRepresentation{
		Id:        ex.Id,
		FileName:  ex.FileName,
		Subject:   ex.Subject,
		Tags:      ex.Tags,
		Level:     ex.Level,
		Author:    ex.AuthorId,
		IsBroken:  ex.IsBroken,
		Class:     ex.Class,
		Position:  ex.Position,
		Mark:      ex.Mark,
		TypeOlymp: ex.TypeOlymp,
		Created:   ex.Created,
	}
}
