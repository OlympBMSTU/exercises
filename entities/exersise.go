package entities

// todo is active хранить инфу о о битом задании
// написать Диме Кузнецову на какую почту что слать
// слать заголовок задание и данные
type ExerciseEntity struct {
	Id       uint
	AuthorId uint
	FileName string
	Tags     []string
	Level    int
	Subject  string
	IsBroken bool
}

func NewExerciseEntity(author uint, filename string, tags []string, level int, subject string) ExerciseEntity {
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

func (entity ExerciseEntity) GetLevel() int {
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

func (enity *ExerciseEntity) GetDataForUpdateEntity(other ExerciseEntity) map[string]interface{} {
	updateMap := make(map[string]interface{}, 0)
	if len(other.FileName) > 0 && other.FileName != enity.FileName {
		updateMap["file_name"] = other.FileName
	}

	// TODO normalize tags to lower
	if len(other.Tags) > 0 {
		if len(enity.Tags) == 0 {
			enity.Tags = other.Tags
			updateMap["tags_to_add"] = enity.Tags
			updateMap["tags_to_remove"] = make([]string, 0)
		} else {
			// copy one to anotther
			to_delete := make([]string, 0)
			new_tags := other.Tags
			for i := 0; i < len(enity.Tags); i += 1 {
				exist := false
				for j := 0; j < len(new_tags); j += 1 {
					if enity.Tags[i] == new_tags[j] {
						exist = true
						// remove equal data from
						new_tags = append(new_tags[:j], new_tags[(j+1):]...)
						break
					}
				}
				if !exist {
					to_delete = append(to_delete, enity.Tags[i])
				}
			}
			updateMap["tags_to_add"] = to_delete
			updateMap["tags_to_remove"] = new_tags
		}
	}

	if other.Level > 0 && other.Level != enity.Level {
		updateMap["level"] = other.Level
	}

	if len(other.Subject) > 0 && other.Subject != enity.Subject {
		updateMap["subject"] = other.Subject
	}

	if other.IsBroken != enity.IsBroken {
		updateMap["is_broken"] = other.IsBroken
	}

	return updateMap
}
