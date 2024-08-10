package utils

import (
	"fmt"
	"log"
	"net/smtp"
)

func SendMail(toMail string, subject string, body string) (err error) {
	email, password, _ := ReadGmailDetails()

	auth := smtp.PlainAuth("", email, password, "smtp.gmail.com")
	to := []string{toMail}

	msg := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s\r\n", toMail, subject, body))

	err = smtp.SendMail("smtp.gmail.com:587", auth, email, to, msg)
	if err != nil {
		log.Printf("could not send mail: %s", err)
		return err
	}

	return nil
}
