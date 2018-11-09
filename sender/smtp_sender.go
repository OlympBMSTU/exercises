package sender

import (
	"bytes"
	"encoding/json"
	"log"
	"net/smtp"

	"github.com/OlympBMSTU/exercises/config"
	"github.com/OlympBMSTU/exercises/logger"
	"github.com/OlympBMSTU/exercises/sender/result"
)

type AnswerS struct {
	ExId   uint
	Answer string
}

// noAuth
type noAuth struct {
}

func (a *noAuth) Start(server *ServerInfo) (string, []byte, error) {
	return "", nil, nil
}

func (a *noAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		// We've already sent everything.
		return nil, errors.New("unexpected server challenge")
	}
	return nil, nil
}

func NoAuth() Auth {
	return &noAuth{}
}

func SendAnswer(exId uint, answer string) result.SenderResult {
	conf, _ := config.GetConfigInstance()
	from := conf.GetSMTPUser()
	pass := conf.GetSMTPPassword()
	to := conf.GetAcceptorMail()
	subject := conf.GetMailSubject()
	path := conf.GetSMTPHost() + ":" + conf.GetSMTPPort()

	answerStruct := AnswerS{exId, answer}
	val, err := json.Marshal(answerStruct)
	if err != nil {
		log.Println(err.Error())
		return result.ErrorResult(err)
	}

	msg := "From: " +
		"Центр довузовской подготовки МГТУ им. Н.Э. Баумана" +
		" <cdp@bmstu.ru>\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n"

	var writer bytes.Buffer
	writer.WriteString(msg)
	writer.Write(val)

	err = smtp.SendMail(path,
		NoAuth(),
		from, []string{to}, writer.Bytes())

	if err != nil {
		logger.LogE.Println(err.Error())
		return result.ErrorResult(err)
	}

	return result.OkResult(nil)
}
