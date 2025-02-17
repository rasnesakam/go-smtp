package main

import (
	"errors"
	"net/smtp"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type MailArgs struct {
	To          []string
	Subject     string
	Content     string
	ContentType string
}

func sendMail(mailArgs MailArgs) error {
	err := godotenv.Load()
	if err != nil {
		return errors.New("error loading .env file")
	}
	from, ok := os.LookupEnv("SMTP_USER")
	if !ok {
		return errors.New("undefined host. Please specify SMTP_USER value")
	}
	smtpHost, ok := os.LookupEnv("SMTP_HOST")
	if !ok {
		return errors.New("undefined host. Please specify SMTP_HOST value")
	}
	smtpPassword, ok := os.LookupEnv("SMTP_PASSWD")
	if !ok {
		return errors.New("undefined password. Please specify SMTP_PASSWD value")
	}
	smtpPort, ok := os.LookupEnv("SMTP_PORT")
	if !ok {
		return errors.New("undefined password. Please specify SMTP_PORT value")
	}

	headers := make(map[string]string)
	headers["From"] = from
	headers["To"] = strings.Join(mailArgs.To, ",") // Alıcı adresini string olarak ekleyin
	headers["Subject"] = mailArgs.Subject
	headers["Content-Type"] = mailArgs.ContentType // İçerik tipini belirtin

	// E-posta içeriğini oluşturma
	message := ""
	for k, v := range headers {
		message += k + ": " + v + "\r\n"
	}
	message += "\r\n" + mailArgs.Content

	auth := smtp.PlainAuth("", from, smtpPassword, smtpHost)
	return smtp.SendMail(smtpHost+":"+smtpPort, auth, from, mailArgs.To, []byte(message))
}
