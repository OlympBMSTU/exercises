package entities

// todo is active хранить инфу о о битом задании
// написать Диме Кузнецову на какую почту что слать
// слать заголовок задание и данные
type ExcercieseEntity struct {
	Id          uint
	AuthorId    uint
	FileName    string
	RightAnswer string
	Tags        []string
	Level       uint
	Subject     string
}

func NewExcercieseEntity(answer string, tags []string, level uint, subject string) ExcercieseEntity {
	return ExcercieseEntity{
		Tags:        tags,
		RightAnswer: answer,
		Level:       level,
		Subject:     subject,
	}
}

func NewExcercieseEntity1(author_id uint, file_name string, right_answer string) ExcercieseEntity {
	return ExcercieseEntity{
		AuthorId:    author_id,
		FileName:    file_name,
		RightAnswer: right_answer,
	}
}

func (entity *ExcercieseEntity) GetId() uint {
	return entity.Id
}

func (entity *ExcercieseEntity) GetAuthorId() uint {
	return entity.AuthorId
}

func (entity *ExcercieseEntity) GetFileName() string {
	return entity.FileName
}

func (entity *ExcercieseEntity) GetRightAnswer() string {
	return entity.RightAnswer
}

func (entity *ExcercieseEntity) GetLevel() uint {
	return entity.Level
}

func (entity *ExcercieseEntity) GetSubject() string {
	return entity.Subject
}

func (entity *ExcercieseEntity) GetTags() []string {
	return entity.Tags
}

func (entity *ExcercieseEntity) SetFileName(file_name string) {
	entity.FileName = file_name
}

func (entity *ExcercieseEntity) SetAuthor(author uint) {
	entity.AuthorId = author
}
