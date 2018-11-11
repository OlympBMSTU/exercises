package sender

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"github.com/OlympBMSTU/exercises/config" //"github.com/OlympBMSTU/exercises/logger"
	"github.com/OlympBMSTU/exercises/sender/result"
	"log"
	"net/smtp"
)

type AnswerS struct {
	Id     uint   `json:"id"`
	Answer string `json:"answer"`
}

func SendAnswer(ExId uint, Answer string) result.SenderResult {
	conf, _ := config.GetConfigInstance()
	//from := conf.GetSMTPUser()
	//pass := conf.GetSMTPPassword()
	to := conf.GetAcceptorMail()
	subject := conf.GetMailSubject()
	//path := conf.GetSMTPHost() + ":" + conf.GetSMTPPort()

	answerStruct := AnswerS{ExId, Answer}
	val, err := json.Marshal(answerStruct)
	if err != nil {
		log.Println(err.Error())
		return result.ErrorResult(err)
	}

	encoded := base64.StdEncoding.EncodeToString(val)

	msg := "From: " +
		"Центр довузовской подготовки МГТУ им. Н.Э. Баумана" +
		" <cdp@bmstu.ru>\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" + encoded

	// Connect to the remote SMTP server.
	log.Println("1")
	c, err := smtp.Dial("mail.bmstu.ru:25")
	if err != nil {
		log.Println(err)
	}
	defer c.Close()
	// Set the sender and recipient.
	log.Println("2")
	c.Mail("cdp@bmstu.ru")
	c.Rcpt("himjune@mail.ru")
	// Send the email body.
	log.Println("3")
	wc, err := c.Data()
	if err != nil {
		log.Println("OnData")
		log.Println(err)
	}
	defer wc.Close()

	log.Println("4")
	buf := bytes.NewBufferString(msg)
	if _, err = buf.WriteTo(wc); err != nil {
		log.Println(err)
	}

	return result.OkResult(nil)
}

/*var writer bytes.Buffer
writer.WriteString(msg)
writer.Write(val)

err = smtp.SendMail(path,
	NoAuth(),
	from, []string{to}, writer.Bytes())

if err != nil {
	logger.LogE.Println(err.Error())
	return result.ErrorResult(err)
}*/
