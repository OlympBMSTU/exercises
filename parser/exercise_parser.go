package parser

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/OlympBMSTU/exercises/parser/result"
	"github.com/OlympBMSTU/exercises/views"
)

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
		return result.ErrorResult(result.INCORRECT_LEVEL, "Level is broken")
	}

	res := views.ExerciseView{
		Answer:   answer,
		FileName: "",
		Subject:  subjectArr[0],
		Level:    uint(level),
		Tags:     tags,
	}

	return result.OkResult(res)
}
