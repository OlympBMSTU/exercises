package sender

import (
	"bytes"
	"encoding/json"
	"net/smtp"

	"github.com/OlympBMSTU/exercises/config"
	"github.com/OlympBMSTU/exercises/sender/result"
)

type AnswerS struct {
	ExId   uint
	Answer string
}

func SendAnswer(exId uint, answer string) result.SenderResult {
	conf, _ := config.GetConfigInstance()
	from := conf.GetSMTPUser()
	pass := conf.GetSMTPPassword()
	to := conf.GetAcceptorMail()
	subject := conf.GetMailSubject()

	answerStruct := AnswerS{exId, answer}
	val, err := json.Marshal(answerStruct)
	if err != nil {
		return result.ErrorResult(err)
	}

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n"

	var writer bytes.Buffer
	writer.WriteString(msg)
	writer.Write(val)

	err = smtp.SendMail("smtp.yandex.ru:25",
		smtp.PlainAuth("", from, pass, "smtp.yandex.ru"),
		from, []string{to}, writer.Bytes())

	if err != nil {
		return result.ErrorResult(err)
	}

	return result.OkResult(nil)
}
