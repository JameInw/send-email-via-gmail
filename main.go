package main

import (
	"log"
	email "send-email-via-gmail/controller"
	"send-email-via-gmail/util"
	"strings"
)

func main() {

	logger := util.NewCustomLogger("Main 🤖")

	// Get Environment YAML Using Base64
	Config, err := util.GetEnvironmentYAML()
	if err != nil {
		log.Println(err)
	}

	logger.Info("Config:", "GenesysCloud",
		strings.Join([]string{
			"🔼 From: " + Config.Email.From,
			"🔑 Password: " + Config.Email.Password,
			"⚒️ SMTPHost: " + Config.Email.SMTPHost,
			"🛠️ SMTPPort: " + Config.Email.SMTPPort,
			"🔽 To: " + strings.Join(Config.Email.To, ","),
			"🔽 Cc: " + strings.Join(Config.Email.Cc, ","),
			"📄 Subject: " + Config.Email.Subject,
		}, "\n"),
	)

	email.NewSendEmail(Config, "./export/test-file.txt")

}
