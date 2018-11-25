package parser

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"github.com/OlympBMSTU/exercises/parser/result"
	"github.com/OlympBMSTU/exercises/views"
)

var gradeArr = []string{"11 класс", "10 класс", "9 класс"}

func ParseExViewPostForm(form map[string][]string) result.ParserResult {
	var err error
	answerArr := form["answer"]
	subjectArr := form["subject"]
	levelStringArr := form["level"]
	tagsJsonArr := form["tags"]

	if len(answerArr) < 1 || len(subjectArr) < 1 || len(levelStringArr) < 1 || len(tagsJsonArr) < 1 {
		return result.ErrorResult(result.INCORRECT_BODY, "Some fileds is empty")
	}

	var rawTags []string
	err = json.Unmarshal([]byte(tagsJsonArr[0]), &rawTags)
	if err != nil {
		log.Println(err.Error())
		return result.ErrorResult(result.INCORRECT_TAGS, "Some tags array is broken")
	}

	tags := make([]string, 0)
	for _, tag := range rawTags {
		trimTag := strings.Trim(tag, " ")
		tags = append(tags, trimTag)
	}

	answer := strings.Trim(answerArr[0], " ")
	answer = strings.Replace(answer, ",", ".", -1)

	var level int
	if level, err = strconv.Atoi(levelStringArr[0]); err != nil {
		log.Println(err.Error())
		return result.ErrorResult(result.INCORRECT_LEVEL, "Level is broken")
	}

	res := views.ExerciseView{
		Answer:   answer,
		FileName: "",
		Subject:  subjectArr[0],
		Level:    level,
		Tags:     tags,
	}

	return result.OkResult(res)

	// var err error
	// answerArr := form["answer"]
	// subjectArr := form["subject"]
	// levelStringArr := form["level"]
	// tagsJSONArr := form["tags"]

	// if len(answerArr) < 1 || len(subjectArr) < 1 || len(levelStringArr) < 1 || len(tagsJSONArr) < 1 {
	// 	return result.ErrorResult(result.INCORRECT_BODY, "Some fileds is empty")
	// }

	// var rawTags []string
	// err = json.Unmarshal([]byte(tagsJSONArr[0]), &rawTags)
	// if err != nil {
	// 	log.Println(err.Error())
	// 	return result.ErrorResult(result.INCORRECT_TAGS, "Some tags array is broken")
	// }

	// // check tags
	// // rawTags[0] - grade, rawTags[1] - rate
	// tags := make([]string, 0)
	// if len(rawTags) != 2 {
	// 	return result.ErrorResult(result.INCORRECT_TAGS, "Count tags is not equal 2")
	// }

	// gradeIsCorrect := false
	// for i, tag := range rawTags {
	// 	// normalize tag
	// 	trimTag := strings.Trim(tag, " ")
	// 	trimTag = strings.ToLower(trimTag)

	// 	// check grade is correct
	// 	if i == 0 {
	// 		for _, grade := range gradeArr {
	// 			if grade == trimTag {
	// 				gradeIsCorrect = true
	// 				break
	// 			}
	// 		}
	// 		if !gradeIsCorrect {
	// 			return result.ErrorResult(result.INCORRECT_TAGS, "Grade is broken")
	// 		}
	// 	}

	// 	// check that rate is number
	// 	if i == 1 {
	// 		_, err := strconv.Atoi(trimTag)
	// 		if err != nil {
	// 			return result.ErrorResult(result.INCORRECT_TAGS, "Rate is broken")
	// 		}
	// 	}
	// 	tags = append(tags, trimTag)
	// }

	// answer := strings.Trim(answerArr[0], " ")
	// answer = strings.Replace(answer, ",", ".", -1)

	// var level int
	// if level, err = strconv.Atoi(levelStringArr[0]); err != nil {
	// 	log.Println(err.Error())
	// 	return result.ErrorResult(result.INCORRECT_LEVEL, "Level is broken")
	// }

	// res := views.ExerciseView{
	// 	Answer:   answer,
	// 	FileName: "",
	// 	Subject:  subjectArr[0],
	// 	Level:    level,
	// 	Tags:     tags,
	// }

	// return result.OkResult(res)
}

func ParseExViewPostUpdateForm(form map[string][]string) result.ParserResult {
	var err error
	idStrArr := form["id"]
	answerArr := form["answer"]
	subjectArr := form["subject"]
	levelStringArr := form["level"]
	tagsJSONArr := form["tags"]
	// file := form["file"]
	// fmt.Println(file)

	id := -1
	if len(idStrArr) == 0 {
		return result.ErrorResult(result.INCORRECT_BODY, "No id presented")
	} else {
		id, err = strconv.Atoi(idStrArr[0])
		if err != nil {
			return result.ErrorResult(result.INCORRECT_BODY, "Id is not a number")
		}
	}

	if len(answerArr) == 0 && len(subjectArr) == 0 && len(levelStringArr) == 0 && len(tagsJSONArr) == 0 {
		return result.OkResult(views.ExerciseView{ID: id})
		//return result.ErrorResult(result.INCORRECT_BODY, "Body has no data")
	}

	tags := make([]string, 0)
	if len(tagsJSONArr) > 0 {
		var rawTags []string
		err = json.Unmarshal([]byte(tagsJSONArr[0]), &rawTags)
		if err != nil {
			log.Println(err.Error())
			return result.ErrorResult(result.INCORRECT_TAGS, "Some tags array is broken")
		}

		for _, tag := range rawTags {
			trimTag := strings.Trim(tag, " ")
			tags = append(tags, trimTag)
		}

		// gradeIsCorrect := false
		// for i, tag := range rawTags {
		// 	// normalize tag
		// 	trimTag := strings.Trim(tag, " ")
		// 	trimTag = strings.ToLower(trimTag)

		// 	// check grade is correct
		// 	if i == 0 {
		// 		for _, grade := range gradeArr {
		// 			if grade == trimTag {
		// 				gradeIsCorrect = true
		// 				break
		// 			}
		// 		}
		// 		if !gradeIsCorrect {
		// 			return result.ErrorResult(result.INCORRECT_TAGS, "Grade is broken")
		// 		}
		// 	}

		// 	// check that rate is number
		// 	if i == 1 {
		// 		_, err := strconv.Atoi(trimTag)
		// 		if err != nil {
		// 			return result.ErrorResult(result.INCORRECT_TAGS, "Rate is broken")
		// 		}
		// 	}
		// 	tags = append(tags, trimTag)
		// }
	}

	answer := ""
	if len(answerArr) > 0 {
		answer := strings.Trim(answerArr[0], " ")
		answer = strings.Replace(answer, ",", ".", -1)
	}

	level := -1
	if len(levelStringArr) > 0 {
		if level, err = strconv.Atoi(levelStringArr[0]); err != nil {
			log.Println(err.Error())
			return result.ErrorResult(result.INCORRECT_LEVEL, "Level is broken")
		}
	}

	res := views.ExerciseView{
		ID:       id,
		Answer:   answer,
		FileName: "",
		Subject:  subjectArr[0],
		Level:    level,
		Tags:     tags,
	}

	return result.OkResult(res)
}
