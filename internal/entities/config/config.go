package config

type Config struct {
	SendingHours []int
	Email        Email
}

type Email struct {
	From     string
	Pass     string
	To       string
	SmtpHost string
	SmtpPort int
}
