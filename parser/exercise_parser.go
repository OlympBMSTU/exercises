package parser

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"github.com/OlympBMSTU/exercises/logger"
	"github.com/OlympBMSTU/exercises/parser/result"
	"github.com/OlympBMSTU/exercises/views"
)

const (
	FIRST_ROUND  = 1
	SECOND_ROUND = 2
)

var gradeArr = []string{"11 класс", "10 класс", "9 класс"}

//

func ParseExViewPostForm(form map[string][]string) result.ParserResult {
	var err error
	answerArr := form["answer"]
	subjectArr := form["subject"]
	levelStringArr := form["level"]
	tagsJSONArr := form["tags"]
	classArr := form["class"]
	positionArr := form["position"]
	markArr := form["mark"]
	typeOlympArr := form["type_olymp"]
	answersArr := form["answers"]
	log := logger.GetLogger()

	// maybe no need

	if len(typeOlympArr) < 1 {
		log.Warn("Type olymp is incorrect", typeOlympArr)
		return result.ErrorResult(result.INCORRECT_TYPE_OLYMP, "Empty type olymp")
	}

	typeOlymp, err := strconv.Atoi(typeOlympArr[0])
	if err != nil {
		log.Warn("Type olymp is incorrect", typeOlympArr)
		return result.ErrorResult(result.INCORRECT_TYPE_OLYMP, "Incorrect type olymp")
	}

	if typeOlymp != FIRST_ROUND && typeOlymp != SECOND_ROUND {
		log.Warn("Type olymp incorrect round", typeOlymp)
		return result.ErrorResult(result.INCORRECT_TYPE_OLYMP, "Incorrect type olymp")
	}

	arrExist := func() bool {
		if typeOlymp == FIRST_ROUND {
			return len(answerArr) > 0
		}
		return len(answersArr) > 0
	}()

	if len(subjectArr) < 1 || len(levelStringArr) < 1 || len(tagsJSONArr) < 1 ||
		len(classArr) < 1 || len(positionArr) < 1 || len(markArr) < 1 || len(typeOlympArr) < 1 || !arrExist {
		log.Error("Some fields in request empty", nil)
		return result.ErrorResult(result.INCORRECT_BODY, "Some fileds is empty")
	}

	var rawTags []string
	err = json.Unmarshal([]byte(tagsJSONArr[0]), &rawTags)
	if err != nil {
		log.Error("Error parse tags", err)
		return result.ErrorResult(result.INCORRECT_TAGS, "Some tags array is broken")
	}

	tags := make([]string, 0)
	for _, tag := range rawTags {
		trimTag := strings.Trim(tag, " ")
		tags = append(tags, trimTag)
	}

	var answers []views.AnswerView
	answer := ""
	if typeOlymp == FIRST_ROUND {
		answer := strings.Trim(answerArr[0], " ")
		answer = strings.Replace(answer, ",", ".", -1)
	} else {
		err = json.Unmarshal([]byte(answersArr[0]), &answers)
		if err != nil {
			// fmt.Print(err.Error()) // To logger
			log.Error("Error parse answers", err)
			return result.ErrorResult(result.INCORRECT_ANSWER_ARR, "Some answers array is broken")
		}
	}

	var level int
	if level, err = strconv.Atoi(levelStringArr[0]); err != nil {
		log.Error("Error parse level", err) //Println(err.Error())
		return result.ErrorResult(result.INCORRECT_LEVEL, "Level is broken")
	}

	var class int
	if class, err = strconv.Atoi(classArr[0]); err != nil {
		log.Error("Error parse class", err)
		return result.ErrorResult(result.INCORRECT_CLASS, "Class is broken")
	}

	var mark int
	if mark, err = strconv.Atoi(classArr[0]); err != nil {
		log.Error("Error parse mark", err)
		return result.ErrorResult(result.INCORRECT_MARK, "Mark is broken")
	}

	var position int
	if position, err = strconv.Atoi(classArr[0]); err != nil {
		log.Error("Error parse position", err)
		return result.ErrorResult(result.INCORRECT_POSITION, "Position is broken")
	}

	res := views.ExerciseView{
		Answer:    answer,
		FileName:  "",
		Subject:   subjectArr[0],
		Level:     level,
		Tags:      tags,
		Class:     class,
		Mark:      mark,
		TypeOlymp: typeOlymp,
		Position:  position,
		Answers:   answers,
	}

	return result.OkResult(res)
}

func ParseExViewPostUpdateForm(form map[string][]string) result.ParserResult {
	var err error
	idStrArr := form["id"]
	answerArr := form["answer"]
	subjectArr := form["subject"]
	levelStringArr := form["level"]
	tagsJSONArr := form["tags"]
	isBrokenArr := form["is_broken"]
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

	if len(answerArr) == 0 && len(subjectArr) == 0 && len(levelStringArr) == 0 && len(tagsJSONArr) == 0 || len(isBrokenArr) == 0 {
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

	subject := ""
	if len(subjectArr) > 0 {
		subject = subjectArr[0]
	}

	isBroken := false
	if len(isBrokenArr) > 0 {
		isBroken = func() bool {
			if isBrokenArr[0] == "true" {
				return true
			} else {
				return false
			}
		}()
	}

	res := views.ExerciseView{
		ID:       id,
		Answer:   answer,
		FileName: "",
		Subject:  subject,
		Level:    level,
		Tags:     tags,
		IsBroken: isBroken,
	}

	return result.OkResult(res)
}
