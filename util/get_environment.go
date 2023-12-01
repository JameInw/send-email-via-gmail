package util

import (
	"encoding/base64"
	"os"
	"send-email-via-gmail/models"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

func GetEnvironmentYAML() (models.Config, error) {
	logger := NewCustomLogger("util ⚒️ | GetEnvironmentYAML")

	var Config models.Config

	yamlFile, err := os.ReadFile("configuration.yaml")
	if err != nil {
		logger.Error("Unable to read YAML file:", "err", err)
	}

	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		logger.Error("Error converting YAML:", "err", err)
	}

	// Decode base64 Email
	DFrom := decodeBase64(Config.Email.From)
	DPassword := decodeBase64(Config.Email.Password)
	Config.Email.From = DFrom
	Config.Email.Password = DPassword

	var DTo []string
	for _, dataTo := range Config.Email.To {
		subDTo := decodeBase64(dataTo)
		DTo = append(DTo, subDTo)
	}
	Config.Email.To = DTo

	var DCc []string
	for _, dataCc := range Config.Email.Cc {
		subDCc := decodeBase64(dataCc)
		DCc = append(DCc, subDCc)
	}
	Config.Email.Cc = DCc

	// Setup Date Time
	currentTime := time.Now()
	formattedemail := currentTime.Format("02-01-2006")
	Config.Email.Subject = strings.Replace(Config.Email.Subject, "{DD-MM-YYYY}", formattedemail, 1)

	logger.Debug("Config:", "GenesysCloud",
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

	return Config, nil
}

func decodeBase64(variable string) string {
	logger := NewCustomLogger("util ⚒️ | decodeBase64")

	Decode, err := base64.StdEncoding.DecodeString(variable)
	if err != nil {
		logger.Error("An error occurred decoding:", "err", err)
	}

	return string(Decode)

}
