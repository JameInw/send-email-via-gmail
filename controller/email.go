package email

import (
	"send-email-via-gmail/models"
	"send-email-via-gmail/util"
	"strconv"

	"gopkg.in/gomail.v2"
)

func NewSendEmail(Config models.Config, Pathfile string) {
	logger := util.NewCustomLogger("controller 🚀 | NewSendEmail")

	body := "This is Test Send Email via Gmail SMTP"

	m := gomail.NewMessage()
	m.SetHeader("From", Config.Email.From)
	m.SetHeader("To", Config.Email.To...)
	m.SetHeader("Cc", Config.Email.Cc...)
	m.SetHeader("Subject", Config.Email.Subject)
	m.SetBody("text/html", body)
	m.Attach(Pathfile)

	SMTPPort, err := strconv.Atoi(Config.Email.SMTPPort)
	if err != nil {
		logger.Error("Error conv Str to int", "err", err)
	}

	d := gomail.NewDialer(Config.Email.SMTPHost, SMTPPort, Config.Email.From, Config.Email.Password)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		logger.Error("Error Dial And Send Email", "err", err)
		panic(err)
	} else {
		logger.Info("📨 Email sent successfully.")
	}

}
