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

// func send(body string) {
// 	conf, _ := Config.GetInstance()
// 	from := conf.GetSenderMail()
// 	pass := conf.GetSenderPass()
// 	to := conf.GetAcceptorMail()

// 	val, err := json.Marshal(data)
// 	if err != nil {
// 		return
// 	}
// 	msg := "From: " + from + "\n" +
// 		"To: " + to + "\n" +
// 		"Subject: Hello there\n\n"
// 	var writer bytes.Buffer
// 	writer.WriteString(msg)
// 	writer.Write(val)
// 	err = smtp.SendMail("smtp.yandex.ru:25",
// 		smtp.PlainAuth("", from, pass, "smtp.yandex.ru"),
// 		from, []string{to}, writer.Bytes())

// 	if err != nil {
// 		log.Printf("smtp error: %s", err)
// 		return
// 	}

// 	// log.Print("sent, visit http://foobarbazz.mailinator.com")
// }

func SendAnswer(ex_id uint, answer string) result.SenderResult {
	conf, _ := config.GetConfigInstance()
	from := conf.GetSMTPUser()
	pass := conf.GetSMTPPassword()
	to := conf.GetAcceptorMail()
	subject := conf.GetMailSubject()

	answerStruct := AnswerS{ex_id, answer}
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
