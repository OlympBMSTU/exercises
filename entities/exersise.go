package entities

// todo is active хранить инфу о о битом задании
// написать Диме Кузнецову на какую почту что слать
// слать заголовок задание и данные
type ExerciseEntity struct {
	Id       uint
	AuthorId uint
	FileName string
	Tags     []string
	Level    uint
	Subject  string
	IsBroken bool
}

func NewExerciseEntity(author uint, filename string, tags []string, level uint, subject string) ExerciseEntity {
	return ExerciseEntity{
		AuthorId: author,
		FileName: filename,
		Tags:     tags,
		Level:    level,
		Subject:  subject,
		IsBroken: false,
	}
}

func (entity ExerciseEntity) GetId() uint {
	return entity.Id
}

func (entity ExerciseEntity) GetAuthorId() uint {
	return entity.AuthorId
}

func (entity ExerciseEntity) GetFileName() string {
	return entity.FileName
}

func (entity ExerciseEntity) GetLevel() uint {
	return entity.Level
}

func (entity ExerciseEntity) GetSubject() string {
	return entity.Subject
}

func (entity ExerciseEntity) GetTags() []string {
	return entity.Tags
}

func (entity *ExerciseEntity) SetFileName(file_name string) {
	entity.FileName = file_name
}

func (entity *ExerciseEntity) SetAuthor(author uint) {
	entity.AuthorId = author
}
