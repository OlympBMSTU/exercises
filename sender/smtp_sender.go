package sender

import (
	"encoding/json"
	"fmt"
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

func SendAnswer(ex_id uint, answer string) error {
	answer_struct := AnswerS{ex_id, answer}
	data, err := json.Marshal(answer_struct)
	fmt.Println(data, err)
	return nil
}
