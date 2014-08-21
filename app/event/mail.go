package event

import (
	"net"
	"strconv"

	"github.com/naoina/kocha"
	"gopkg.in/gomail.v1"
)

func SendContactMail(app *kocha.Application, args ...interface{}) error {
	msg := gomail.NewMessage(gomail.SetEncoding(gomail.Base64))
	msg.SetHeader("From", "info@vocanew.kuune.org")
	msg.SetHeader("To", "vocanew@gmail.com")
	msg.SetHeader("Subject", "【ぼかにゅー】ご意見・ご要望")
	msg.SetBody("text/plain", args[0].(string))
	host, ps, err := net.SplitHostPort(kocha.Getenv("VOCANEW_SMTP_ADDR", "localhost:1025"))
	if err != nil {
		return err
	}
	port, err := strconv.Atoi(ps)
	if err != nil {
		return err
	}
	mailer := gomail.NewMailer(host, "", "", port)
	return mailer.Send(msg)
}
