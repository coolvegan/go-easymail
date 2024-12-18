package main

import (
	"log"
	"os"
	"sync"
	"time"
)

func main() {
	//Just testing
	var (
		RecipientName = os.Getenv("RECNAME")
		RecpientEmail = os.Getenv("RECMAIL")
		Subject       = os.Getenv("SUBJECT")
		Message       = os.Getenv("MESSAGE")
	)
	var e Emailer
	var wg sync.WaitGroup
	e = NewLazymail(os.Getenv("SMTPSERVER"), os.Getenv("SMTPUSER"), os.Getenv("SMTPPASSWORD"), BasicEmailSend(1000))
	// e = NewEmail(os.Getenv("SMTPSERVER"), os.Getenv("SMTPUSER"), os.Getenv("SMTPPASSWORD"))
	wg.Add(1)
	go func() {
		i := 0
		f := 0
		for {
			i++
			f++
			err := e.Send(RecipientName, RecpientEmail, Subject, Message)
			if err != nil {

				log.Println(err)
			}
			log.Println("Sende in Main")
			time.Sleep(10 * time.Millisecond)
			if f >= 100 {

				f = 0
				log.Println("8 Sekunden Sendepause")
				time.Sleep(20000 * time.Millisecond)
				wg.Done()
			}
		}
	}()
	wg.Wait()
}
