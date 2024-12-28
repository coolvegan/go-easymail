package eazymail

import (
	"fmt"
	"log"

	"github.com/wneessen/go-mail"
)

type Emailer interface {
	Send(sender, recipient, subject, body string) error
	newEmail() *mail.Client
}

type Email struct {
	username   string
	password   string
	smtpserver string
}

func (e *Email) Send(sender, recipient, subject, body string) error {
	if len(sender) == 0 || len(recipient) == 0 || len(subject) == 0 {
		return fmt.Errorf("Sender: %s, Recepient: %s, Subject: %s", sender, recipient, body)
	}
	message := mail.NewMsg()
	if err := message.FromFormat(sender, e.username); err != nil {
		log.Printf("failed to set FROM address: %s", err)
	}

	if err := message.To(recipient); err != nil {
		log.Printf("failed to set TO address: %s", err)
	}
	message.Subject(subject)
	message.SetBodyString(mail.TypeTextHTML, body)

	client := e.newEmail()

	if err := client.DialAndSend(message); err != nil {
		log.Printf("failed to deliver mail: %s", err)
		return err
	}
	return nil
}

func (e *Email) newEmail() *mail.Client {
	if len(e.smtpserver) == 0 || len(e.username) == 0 || len(e.password) == 0 {
		log.Printf("Username: %s Smtpserver: %s Password: %s missing value(s).", e.username, e.smtpserver, e.password)
	}
	client, err := mail.NewClient(e.smtpserver,
		mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithTLSPortPolicy(mail.TLSMandatory),
		mail.WithUsername(e.username), mail.WithPassword(e.password))
	if err != nil {
		log.Printf("Email Error: %s", err)
	}
	return client
}

func NewEmail(smtpserver, username, password string) *Email {
	return &Email{smtpserver: smtpserver, username: username, password: password}
}
