package eazymail

import (
	"fmt"
	"log"
	"time"

	"github.com/wneessen/go-mail"
)

type LazyMail struct {
	email         Email
	messageQueue  ConcurrentQueue
	BackgroundJob func(l *LazyMail)
}

func (l *LazyMail) Send(sender, recipient, subject, body string) error {
	(*l).messageQueue.enqueue(Message{sender: sender, recipient: recipient, subject: subject, body: body})
	return nil
}

func (l *LazyMail) newEmail() *mail.Client {
	return l.email.newEmail()
}

func (l *LazyMail) run() {
	l.BackgroundJob(l)
}

func lazyTest(seconds int) func(l *LazyMail) {
	return func(l *LazyMail) {
		for {
			time.Sleep(time.Duration(seconds) * time.Millisecond)
			msg, err := l.messageQueue.dequeue()
			if err != nil {
				continue
			}
			fmt.Printf("\n Empfänger: %s\n Betreff: %s\n Inhalt: %s\n", msg.recipient, msg.subject, msg.body)
		}
	}
}

func BasicEmailSend(seconds int) func(l *LazyMail) {
	return func(l *LazyMail) {
		for {
			time.Sleep(time.Duration(seconds) * time.Millisecond)
			msg, err := l.messageQueue.dequeue()
			if err != nil {
				continue
			}
			go l.email.Send(msg.sender, msg.recipient, msg.subject, msg.body)

		}
	}
}

func SendUpToEightEmailsAndThenDelay(seconds int) func(l *LazyMail) {
	return func(l *LazyMail) {
		i := 0
		for {
			i++
			if i > 8 {
				time.Sleep(time.Duration(seconds) * time.Millisecond)
				i = 0
			}
			msg, err := l.messageQueue.dequeue()
			if err != nil {
				continue
			}
			go l.email.Send(msg.sender, msg.recipient, msg.subject, msg.body)

		}
	}
}

func lazyUnitTest(seconds int, want *[]string) func(l *LazyMail) {
	return func(l *LazyMail) {
		for {
			time.Sleep(time.Duration(seconds) * time.Millisecond)
			msg, err := l.messageQueue.dequeue()
			if err != nil {
				continue
			}
			m := fmt.Sprintf("Empfänger: %s Betreff: %s Inhalt: %s", msg.recipient, msg.subject, msg.body)
			*want = append(*want, m)
		}
	}
}

func NewLazymail(smtpserver, username, password string, behavior func(l *LazyMail)) *LazyMail {
	if len(smtpserver) == 0 || len(username) == 0 || len(password) == 0 {
		log.Panicf("NewLazyMail Error\n SMTPSERVER: %s\n SMTPUSER: %s\n SMTPPASSWORD: %s\n Error: at least one value is empty", smtpserver, username, password)

	}
	l := LazyMail{email: *NewEmail(smtpserver, username, password), BackgroundJob: behavior}
	go l.run()
	return &l
}
