package main

import (
	"log"
	email "send-email-via-gmail/controller"
	"send-email-via-gmail/util"
	"strings"

	"github.com/robfig/cron"
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
		"Worker",
		strings.Join([]string{
			"⌛️ crontab expression : " + Config.Worker.CronTab,
		}, "\n"),
	)

	c := cron.New()

	logger.Info("Run Job ⏳", "Crontab", Config.Worker.CronTab)

	c.AddFunc(Config.Worker.CronTab, func() {
		email.NewSendEmail(Config, "./export/test-file.txt")
	})

	c.Start()
	select {} // Wait for the cron job to run (the program can be run continuously)

}
