package utils

import (
	"fmt"
	"net/smtp"
)

func SendResetEmail(email, token string) error {
	from := "your-email@example.com"
	password := "your-email-password"
	to := email
	smtpHost := "smtp.example.com"
	smtpPort := "587"

	message := fmt.Sprintf("Subject: Password Reset\n\nClick the link to reset your password: https://example.com/reset?token=%s", token)

	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(message))
	if err != nil {
		return err
	}

	return nil
}
