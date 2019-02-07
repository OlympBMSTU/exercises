package entities

import "strings"

// todo is active хранить инфу о о битом задании
// написать Диме Кузнецову на какую почту что слать
// слать заголовок задание и данные
type ExerciseEntity struct {
	Id        uint
	AuthorId  uint
	FileName  string
	Tags      []string
	Level     int
	Subject   string
	IsBroken  bool
	Class     int
	Position  int
	Mark      int
	TypeOlymp int
	Answers   []Answer
	Created   string
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

func (entity *ExerciseEntity) GetClass() int {
	return entity.Class
}

func (entity *ExerciseEntity) GetPosition() int {
	return entity.Position
}

func (entity *ExerciseEntity) GetMark() int {
	return entity.Mark
}

func (entity *ExerciseEntity) GetTypeOlymp() int {
	return entity.TypeOlymp
}

func (entity *ExerciseEntity) GetAnswers() []Answer {
	return entity.Answers
}

func (enity *ExerciseEntity) GetDataForUpdateEntity(other ExerciseEntity) map[string]interface{} {
	updateMap := make(map[string]interface{}, 0)
	if len(other.FileName) > 0 && other.FileName != enity.FileName {
		updateMap["file_name"] = other.FileName
	}

	// TODO normalize tags to lower

	if len(other.Subject) > 0 && other.Subject != enity.Subject {
		updateMap["subject"] = other.Subject
		if len(other.Tags) > 0 {
			updateMap["tags"] = other.Tags
			updateMap["tags_to_add"] = other.Tags
			updateMap["tags_to_remove"] = enity.Tags
		} else {
			updateMap["tags"] = enity.Tags
			updateMap["tags_to_add"] = enity.Tags
			updateMap["tags_to_remove"] = enity.Tags
		}
	} else {
		if len(other.Tags) > 0 {
			// What to do chekc all tags for its order
			if len(enity.Tags) == 0 {
				enity.Tags = other.Tags
				updateMap["tags"] = other.Tags
				updateMap["tags_to_add"] = enity.Tags
				updateMap["tags_to_remove"] = make([]string, 0)
			} else {
				tagsUpdated := false
				arrLen := len(enity.Tags)
				if arrLen == len(other.Tags) {
					for i := 0; i < arrLen && !tagsUpdated; i++ {
						if strings.ToLower(enity.Tags[i]) != strings.ToLower(other.Tags[i]) {
							tagsUpdated = true
							break
						}
					}
				} else {
					tagsUpdated = true
				}

				if tagsUpdated {
					updateMap["tags"] = other.Tags
					// copy one to anotther
					toDelete := make([]string, 0)
					// var newTags []string
					// newTags := other.Tags
					newTags := make([]string, len(other.Tags))
					copy(newTags, other.Tags)
					for i := 0; i < len(enity.Tags); i += 1 {
						exist := false
						for j := 0; j < len(newTags); j += 1 {
							if strings.ToLower(enity.Tags[i]) == strings.ToLower(newTags[j]) {
								exist = true
								// remove equal data from
								newTags = append(newTags[:j], newTags[(j+1):]...)
								break
							}
						}
						if !exist {
							toDelete = append(toDelete, enity.Tags[i])
						}
					}
					if len(newTags) > 0 || len(toDelete) > 0 {
						updateMap["tags_to_add"] = newTags
						updateMap["tags_to_remove"] = toDelete
					}
				}
			}
		}
	}

	if other.Level > 0 && other.Level != enity.Level {
		updateMap["level"] = other.Level
	}

	if other.IsBroken != enity.IsBroken {
		updateMap["is_broken"] = other.IsBroken
	}

	return updateMap
}
