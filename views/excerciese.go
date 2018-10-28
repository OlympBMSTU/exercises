package views

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/OlympBMSTU/exercises/views/result"
)

type ExerciseView struct {
	Answer   string   `json:"answer"`
	FileName string   `json:"file_name"`
	Subject  string   `json:"subject"`
	Tags     []string `json:"tags"`
	Level    uint     `json:"level"`
	Author   uint
}

func ExcercieseViewFromForm(form map[string][]string) result.ParserResult {
	var err error
	answer := form["answer"]
	subject := form["subject"]
	level_str := form["level"]
	tags_arr := form["tags"]

	if len(answer) < 1 || len(subject) < 1 || len(level_str) < 1 || len(tags_arr) < 1 {
		return result.ErrorResult(result.INCORRECT_BODY, "Some fileds is empty")
	}

	var raw_tags []string
	err = json.Unmarshal([]byte(tags_arr[0]), &raw_tags)
	if err != nil {
		return result.ErrorResult(result.INCORRECT_TAGS, "Some tags array is broken")
	}

	tags := make([]string, 1)
	for _, tag := range tags {
		trimTag := strings.Trim(tag, " ")
		tags = append(tags, trimTag)
	}

	answer = strings.Trim(answer, " ")
	answer = strings.Replace(answer, ",", ".", -1)

	var level int
	if level, err = strconv.Atoi(level_str[0]); err != nil {
		return result.ErrorResult(result.INCORRECT_LEVEL, "Level is broken")
	}

	res := ExerciseView{
		Answer:   answer[0],
		FileName: "",
		Subject:  subject[0],
		Level:    uint(level),
		Tags:     tags,
	}

	return result.OkResult(res)
}
