package models

type Email struct {
	From     string   `yaml:"from"`
	Password string   `yaml:"password"`
	SMTPHost string   `yaml:"smtphost"`
	SMTPPort string   `yaml:"smtpport"`
	To       []string `yaml:"to"`
	Cc       []string `yaml:"cc"`
	Subject  string   `yaml:"subject"`
}

type Worker struct {
	CronTab string `yaml:"crontab"`
}

type Config struct {
	Email  Email  `yaml:"email"`
	Worker Worker `yaml:"worker"`
}
