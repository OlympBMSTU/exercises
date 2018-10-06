package entities

type ExcercieseEntity struct {
	id           uint
	author_id    uint
	file_name    string
	right_answer string
	tags         []string
	level        uint
}

func NewExcercieseEntity(author_id uint, file_name string, right_answer string) ExcercieseEntity {
	return ExcercieseEntity{
		author_id:    author_id,
		file_name:    file_name,
		right_answer: right_answer,
	}
}

func (entity *ExcercieseEntity) GetAuthorId() uint {
	return entity.author_id
}

func (entity *ExcercieseEntity) GetFileName() string {
	return entity.file_name
}

func (entity *ExcercieseEntity) GetRightAnswer() string {
	return entity.right_answer
}

func (entity *ExcercieseEntity) GetLevel() uint {
	return entity.level
}

func (entity *ExcercieseEntity) SetFileName(file_name string) {
	entity.file_name = file_name
}
