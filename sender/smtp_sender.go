package sender

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"mime/quotedprintable"
	"net/smtp"
	"strings"

	"github.com/OlympBMSTU/exercises/sender/result"
)

// $connect = fsockopen ('mail.bmstu.ru', 25, $errno, $errstr, 30);
// echo fgets($connect);
// $i = 0;
// foreach ($list as $p){
// $i++;
// echo "\n\n----------\n{$p['email']}\n";
// fputs($connect, "HELO 195.19.*.*\r\n");
// echo fgets($connect);
// fputs($connect, "MAIL FROM: cdp@bmstu.ru\n");
// echo fgets($connect);
// fputs($connect, "RCPT TO: ".$p['email']."\n");
// $res = fgets($connect);
// echo $res;
// if ($res[0] == '2') {
// fputs($connect, "DATA\r\n");
// echo fgets($connect);
// fputs($connect, "Content-Type: text/plain; charset=utf-8\n");
// fputs($connect, "From: =?utf-8?B?".base64_encode('Кто-то там из МГТУ им. Н. Э. Баумана')."?= <cdp@bmstu.ru>\n");
// fputs($connect, "To: =?utf-8?B?".base64_encode($p['name'])."?= <".$p['email'].">\n");
// fputs($connect, "Subject: =?utf-8?B?".base64_encode($subject)."?=\n");
// fputs($connect, "\n\n");
// fputs($connect, $p['text']."\r\n");
// fputs($connect, ".\r\n");
// echo "\t\t!!".$p['email'].' '.substr(fgets($connect), 0, 3)."\n";
// }
// fputs($connect, "RSET\r\n");
// echo fgets($connect);
// $sql = "UPDATE letter SET datetime=NOW() WHERE id = {$p['id']}";
// $result = pg_query($sql);
// //if ($i >= 100) break;
// }
// echo '$i='.$i."\n\n";
// fputs($connect, "QUIT\r\n");
// fclose($connect);

type AnswerS struct {
	ExId   uint
	Answer string
}

func send(body string) {
	conf, _ := Config.GetInstance()
	from := conf.GetSenderMail()
	pass := conf.GetSenderPass()
	to := conf.GetAcceptorMail()

	val, err := json.Marshal(data)
	if err != nil {
		return
	}
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there\n\n"
	var writer bytes.Buffer
	writer.WriteString(msg)
	writer.Write(val)
	err = smtp.SendMail("smtp.yandex.ru:25",
		smtp.PlainAuth("", from, pass, "smtp.yandex.ru"),
		from, []string{to}, writer.Bytes())

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	// log.Print("sent, visit http://foobarbazz.mailinator.com")
}

func SendAnswer(ex_id uint, answer string) SmtpResult {
	answerStruct := AnswerS{ex_id, answer}
	data, err := json.Marshal(answer_struct)
	if err != nil {
		return result.ErrorResult(err)
	}

	fmt.Println(data, err)
	return nil
}

const (
	SMTPServer = "smtp.gmail.com"
)

type Sender struct {
	User     string
	Password string
}

func NewSender(Username, Password string) Sender {
	return Sender{Username, Password}
}

func (sender Sender) SendMail(Dest []string, Subject, bodyMessage string) {

	msg := "From: " + sender.User + "\n" +
		"To: " + strings.Join(Dest, ",") + "\n" +
		"Subject: " + Subject + "\n" + bodyMessage

	err := smtp.SendMail(SMTPServer+":587",
		smtp.PlainAuth("", sender.User, sender.Password, SMTPServer),
		sender.User, Dest, []byte(msg))

	if err != nil {
		return result.ErrorResult(err)
	}
}

func (sender Sender) WriteEmail(dest []string, contentType, subject, bodyMessage string) string {

	header := make(map[string]string)
	header["From"] = sender.User

	receipient := ""

	for _, user := range dest {
		receipient = receipient + user
	}

	header["To"] = receipient
	header["Subject"] = subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = fmt.Sprintf("%s; charset=\"utf-8\"", contentType)
	header["Content-Transfer-Encoding"] = "quoted-printable"
	header["Content-Disposition"] = "inline"

	message := ""

	for key, value := range header {
		message += fmt.Sprintf("%s: %s\r\n", key, value)
	}

	var encodedMessage bytes.Buffer

	finalMessage := quotedprintable.NewWriter(&encodedMessage)
	finalMessage.Write([]byte(bodyMessage))
	finalMessage.Close()

	message += "\r\n" + encodedMessage.String()

	return message
}

func (sender *Sender) WriteHTMLEmail(dest []string, subject, bodyMessage string) string {

	return sender.WriteEmail(dest, "text/html", subject, bodyMessage)
}

func (sender *Sender) WritePlainEmail(dest []string, subject, bodyMessage string) string {

	return sender.WriteEmail(dest, "text/plain", subject, bodyMessage)
}
